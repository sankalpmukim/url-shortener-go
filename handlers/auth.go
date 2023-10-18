package handlers

import (
	"html/template"
	"net/http"

	"github.com/sankalpmukim/url-shortener-go/controllers"
	"github.com/sankalpmukim/url-shortener-go/flash"
)

// GET /login
func GetLogin(w http.ResponseWriter, r *http.Request) {
	flashes, err := flash.GetFlash(w, r, "error")
	if err != nil {
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/auth/login.html")
	if err != nil {
		w.Write([]byte("Error"))
	}
	tmpl.Execute(w, flashes)
}

// POST /login
func PostLogin(w http.ResponseWriter, r *http.Request) {
	// retrieve the email and password values
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// Read values from the form data
	email := r.FormValue("email")
	password := r.FormValue("password")

	// check if the email and password combination is valid
	isValid := controllers.CheckUserCredentials(email, password)

	if !isValid {
		// flash error message
		flash.SetFlash(w, "error", []byte("Invalid credentials"))

		// redirect the user to the login page
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}
	encodedPayload := controllers.CreateSecretPayload(email)

	// create a new cookie
	cookie := http.Cookie{
		Name:     "auth",
		Value:    encodedPayload,
		Path:     "/",                      // makes sure it's available for the whole domain
		Domain:   "",                       // leave it empty to default to the domain of the calling script
		MaxAge:   3600,                     // 1 hour in seconds, 0 means no 'Max-Age' attribute set. If negative, delete cookie now.
		Secure:   false,                    // true if you only want to send the cookie over HTTPS
		HttpOnly: true,                     // true if you want to prevent JavaScript access to the cookie
		SameSite: http.SameSiteDefaultMode, // or http.SameSiteLaxMode, http.SameSiteStrictMode, http.SameSiteNoneMode
	}

	// set the cookie
	http.SetCookie(w, &cookie)

	// redirect the user to the homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// GET /signup
func GetSignup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup"))
}

// POST /signup
func PostSignup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup Post"))
}

// GET /logout
func GetLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
