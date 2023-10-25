package handlers

import (
	"net/http"
	"os"

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
