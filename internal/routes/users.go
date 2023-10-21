package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sankalpmukim/url-shortener-go/internal/handlers"
)

var Users *chi.Mux

func init() {
	Users = chi.NewRouter()
	Users.Get("/", handlers.GetUsers)
}
