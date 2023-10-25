package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sankalpmukim/url-shortener-go/internal/controllers"
	"github.com/sankalpmukim/url-shortener-go/internal/cookies"
	"github.com/sankalpmukim/url-shortener-go/internal/database"
	"github.com/sankalpmukim/url-shortener-go/internal/lib"
	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

type dbLinkWithOrigin struct {
	database.Link
	CleanedEndpoint string
	CleanedTarget   string
	Origin          string
}

type GetLinksResponse struct {
	Links     []dbLinkWithOrigin `json:"links"`
	FlashInfo cookies.FlashInfo  `json:"flash_info"`
}

func GetLinks(w http.ResponseWriter, r *http.Request) {
	flashInfo, err := cookies.GetFlashInfo(w, r)
	if err != nil {
		logs.Error("Failed to parse form(flash cookie)", err)
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return
	}
	links, err := database.DB.GetLinks()
	if err != nil {
		logs.Error("Error getting links", err)
		w.Write([]byte("Error"))
		return
	}

	logs.Info("Link", links)

	tmpl, err := lib.FlashTemplates("pkg/templates/links/index.html")
	if err != nil {
		logs.Error("Failed to parse form(flash cookie)", err)
		w.Write([]byte("Error"))
	}
	linksOrigin := make([]dbLinkWithOrigin, len(links))
	for i, link := range links {
		linksOrigin[i] = dbLinkWithOrigin{
			Link:            link,
			CleanedEndpoint: lib.TrimProtocol(os.Getenv("ORIGIN")) + "/" + link.Endpoint,
			CleanedTarget:   lib.TrimProtocol(link.Target),
			Origin:          os.Getenv("ORIGIN"),
		}
	}
	getLinksResponse := GetLinksResponse{
		Links:     linksOrigin,
		FlashInfo: flashInfo,
	}
	tmpl.Execute(w, getLinksResponse)
}

func RedirectLink(w http.ResponseWriter, r *http.Request) {
	endpoint := chi.URLParam(r, "endpoint")
	if endpoint == "" {
		logs.Error("Endpoint not provided")
		http.Error(w, "Endpoint not provided", http.StatusBadRequest)
		return
	}
	link, err := database.DB.GetLinkByEndpoint(endpoint)
	if err != nil {
		logs.Error("Error getting link", err)
		// return 404 if link not found
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	if link.Endpoint == "" {
		logs.Error("Link not found")
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, link.Target, http.StatusSeeOther)
}

type GetEditLinkResponse struct {
	Link      database.Link     `json:"link"`
	FlashInfo cookies.FlashInfo `json:"flash_info"`
}

func GetEditLink(w http.ResponseWriter, r *http.Request) {
	endpoint := chi.URLParam(r, "endpoint")
	if endpoint == "" {
		logs.Error("Endpoint not provided")
		http.Error(w, "Endpoint not provided", http.StatusBadRequest)
		return
	}
	link, err := database.DB.GetLinkByEndpoint(endpoint)
	if err != nil {
		logs.Error("Error getting link", err)
		w.Write([]byte("Error"))
		return
	}
	if link.Endpoint == "" {
		logs.Error("Link not found")
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	tmpl, err := lib.FlashTemplates("pkg/templates/links/edit.html")
	if err != nil {
		logs.Error("Failed to parse form(flash cookie)", err)
		w.Write([]byte("Error"))
	}

	flashInfo, err := cookies.GetFlashInfo(w, r)
	if err != nil {
		logs.Error("Failed to parse form(flash cookie)", err)
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return
	}
	linkWithFlashInfo := GetEditLinkResponse{
		Link:      link,
		FlashInfo: flashInfo,
	}
	tmpl.Execute(w, linkWithFlashInfo)
}

func PostEditLink(w http.ResponseWriter, r *http.Request) {
	oldEndpoint := chi.URLParam(r, "endpoint")
	if oldEndpoint == "" {
		logs.Error("Endpoint not provided")
		http.Error(w, "Endpoint not provided", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		logs.Error("Failed to parse form", err)
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// Read values from the form data
	endpoint := r.FormValue("endpoint")
	target := r.FormValue("target")

	err := database.DB.UpdateLink(oldEndpoint, endpoint, target)
	if err != nil {
		logs.Error("Failed to update link", err)
		http.Error(w, "Failed to update link", http.StatusInternalServerError)
		return
	}

	// flash success message
	cookies.CreateOrAppendFlash(w, r, cookies.SUCCESS, "Link updated successfully")

	// redirect the user to the list links page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetCreateLink(w http.ResponseWriter, r *http.Request) {
	flashInfo, err := cookies.GetFlashInfo(w, r)
	if err != nil {
		logs.Error("Failed to parse form(flash cookie)", err)
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return
	}
	tmpl, err := lib.FlashTemplates("pkg/templates/links/create.html")
	if err != nil {
		logs.Error("Failed to parse form(flash cookie)", err)
		w.Write([]byte("Error"))
	}
	tmpl.Execute(w, flashInfo)
}

func PostCreateLink(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logs.Error("Failed to parse form", err)
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// Read values from the form data
	endpoint := r.FormValue("endpoint")
	target := r.FormValue("target")

	// check if the endpoint already exists
	if database.DB.LinkExists(endpoint) {
		// flash error message
		cookies.CreateOrAppendFlash(w, r, cookies.ERROR, "Endpoint already exists")

		// redirect the user to the create link page
		http.Redirect(w, r, "/links/create", http.StatusSeeOther)
		return
	}

	email := controllers.GetMailFromContext(r)
	user, err := database.DB.GetUserByEmail(email)
	if err != nil {
		logs.Error("Failed to get user", err)
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// create a link
	err = database.DB.CreateLink(database.CreateLink{
		Endpoint:  endpoint,
		Target:    target,
		CreatedBy: user.ID,
	})
	if err != nil {
		logs.Error("Failed to create link", err)
		http.Error(w, "Failed to create link", http.StatusInternalServerError)
		return
	}

	// flash success message
	cookies.CreateOrAppendFlash(w, r, cookies.SUCCESS, "Link created successfully")

	// redirect the user to the list links page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
