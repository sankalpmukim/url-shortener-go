package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sankalpmukim/url-shortener-go/handlers"
)

var Auth *chi.Mux

func init() {
	Auth = chi.NewRouter()
	Auth.Get("/login", handlers.GetLogin)
	Auth.Post("/login", handlers.PostLogin)
	Auth.Get("/signup", handlers.GetSignup)
	Auth.Post("/signup", handlers.PostSignup)
	Auth.Get("/logout", handlers.GetLogout)
}
