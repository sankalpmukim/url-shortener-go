package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sankalpmukim/url-shortener-go/internal/handlers"
)

var Links *chi.Mux

func init() {
	Links = chi.NewRouter()
	Links.Get("/", handlers.GetLinks)
}
