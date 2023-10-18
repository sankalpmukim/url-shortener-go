package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/sankalpmukim/url-shortener-go/middleware"
	"github.com/sankalpmukim/url-shortener-go/routes"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello World!"))
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.Write([]byte("Error"))
	}
	tmpl.Execute(w, "Hello from code")
}

func main() {
	initializeEnv()
	r := chi.NewRouter()
	r.Use(chimiddleware.Logger)
	r.Mount("/auth", routes.Auth)
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticated)
		r.Get("/", handleIndex)
	})
	http.ListenAndServe(":3000", r)
}

func initializeEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
