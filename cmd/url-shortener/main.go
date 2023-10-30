package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	i "github.com/sankalpmukim/url-shortener-go/internal/initialize"
	"github.com/sankalpmukim/url-shortener-go/internal/middleware"
	"github.com/sankalpmukim/url-shortener-go/internal/routes"
	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

func main() {
	if err := i.InitAll(); err != nil {
		logs.Error("Error occured during initialization", err)
		panic(err)
	}

	// Get the value of the DEBUG environment variable
	DEBUG := os.Getenv("DEBUG")
	if DEBUG != "true" {
		// cannot use logs package here because it
		// doesn't print to the console.
		fmt.Printf("DEBUG = %v\n", DEBUG)
	}

	// Create a new base router
	r := chi.NewRouter()
	r.Use(chimiddleware.Logger)
	r.Mount("/auth", routes.Auth)

	// Debug only routes
	if DEBUG == "true" {
		r.Mount("/users", routes.Users)
	}

	// Auth Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticated)
		r.Mount("/", routes.Links)
	})

	// Listen on port 3000
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	logs.Info("Starting server on port " + PORT)
	http.ListenAndServe(":"+PORT, r)
}
