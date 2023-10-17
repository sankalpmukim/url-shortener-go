package handlers

import (
	"net/http"
)

// GET /login
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

// POST /login
func HandleLoginPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Post"))
}

// GET /signup
func HandleSignup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup"))
}

// POST /signup
func HandleSignupPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup Post"))
}

// GET /logout
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
