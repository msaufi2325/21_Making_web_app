package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/msaufi2325/21_Making_web_app/pkg/config"
	"github.com/msaufi2325/21_Making_web_app/pkg/handlers"
	"github.com/msaufi2325/21_Making_web_app/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	http.ListenAndServe(portNumber, nil)
}
