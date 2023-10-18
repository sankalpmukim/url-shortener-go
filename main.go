package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handleIndex)
	http.ListenAndServe(":3000", r)
}
