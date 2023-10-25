package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sankalpmukim/url-shortener-go/internal/handlers"
)

var Links *chi.Mux

func init() {
	Links = chi.NewRouter()
	Links.Get("/", handlers.GetLinks)
	Links.Get("/{endpoint}", handlers.RedirectLink)
	Links.Get("/create", handlers.GetCreateLink)
	Links.Post("/create", handlers.PostCreateLink)
	Links.Get("/{endpoint}/edit", handlers.GetEditLink)
	Links.Post("/{endpoint}/edit", handlers.PostEditLink)
}
