package handlers

import (
	"html/template"
	"net/http"

	"github.com/sankalpmukim/url-shortener-go/controllers"
	"github.com/sankalpmukim/url-shortener-go/cookies"
	"github.com/sankalpmukim/url-shortener-go/logs"
)

// GET /login
func GetLogin(w http.ResponseWriter, r *http.Request) {
	flashes, err := cookies.GetFlash(w, r, "error")
	if err != nil {
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/auth/login.html")
	if err != nil {
		w.Write([]byte("Error"))
	}
	tmpl.Execute(w, string(flashes))
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
		cookies.SetFlash(w, "error", []byte("Invalid credentials"))

		// redirect the user to the login page
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}
	encodedPayload := controllers.CreateSecretPayload(email)

	// create a new cookie
	cookie := cookies.CreateCookie("auth", encodedPayload)

	// set the cookie
	http.SetCookie(w, &cookie)

	// redirect the user to the homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// GET /signup
func GetSignup(w http.ResponseWriter, r *http.Request) {
	flashes, err := cookies.GetFlash(w, r, "error")
	if err != nil {
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/auth/signup.html")
	if err != nil {
		w.Write([]byte("Error"))
	}
	tmpl.Execute(w, string(flashes))
}

// POST /signup
func PostSignup(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logs.Error("Failed to parse form", err)
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// Read values from the form data
	fullName := r.FormValue("full_name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	// check if password and confirm password match
	if password != confirmPassword {
		// flash error message
		cookies.SetFlash(w, "error", []byte("Passwords do not match"))

		// redirect the user to the signup page
		http.Redirect(w, r, "/auth/signup", http.StatusSeeOther)
		return
	}

	// check if the email already exists
	if controllers.CheckIfEmailExists(email) {
		// flash error message
		cookies.SetFlash(w, "error", []byte("Email already exists"))

		// redirect the user to the signup page
		http.Redirect(w, r, "/auth/signup", http.StatusSeeOther)
		return
	}

	// create a user
	_, err := controllers.CreateUser(fullName, email, password)

	if err != nil {
		logs.Error("Failed to create user", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// redirect to the login page, with flash message
	cookies.SetFlash(w, "error", []byte("User created successfully"))
	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

// GET /logout
func GetLogout(w http.ResponseWriter, r *http.Request) {
	flashes, err := cookies.GetFlash(w, r, "error")
	if err != nil {
		http.Error(w, "Failed to parse form(flash cookie)", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/auth/logout.html")
	if err != nil {
		w.Write([]byte("Error"))
	}
	// clear "auth" cookie
	dc := cookies.DeleteCookieCookie("auth")

	// set the cookie
	http.SetCookie(w, dc)
	tmpl.Execute(w, string(flashes))
}
