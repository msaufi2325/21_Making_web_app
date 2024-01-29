package handlers

import (
	"net/http"

	"github.com/msaufi2325/21_Making_web_app/pkg/config"
	"github.com/msaufi2325/21_Making_web_app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}
