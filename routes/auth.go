package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sankalpmukim/url-shortener-go/handlers"
)

var Auth *chi.Mux

func init() {
	Auth = chi.NewRouter()
	Auth.Get("/login", handlers.HandleLogin)
	Auth.Post("/login", handlers.HandleLoginPost)
	Auth.Get("/signup", handlers.HandleSignup)
	Auth.Post("/signup", handlers.HandleSignupPost)
	Auth.Get("/logout", handlers.HandleLogout)
}
