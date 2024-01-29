package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/msaufi2325/21_Making_web_app/pkg/config"
	"github.com/msaufi2325/21_Making_web_app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
