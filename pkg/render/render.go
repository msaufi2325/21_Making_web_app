package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// RenderTemplate render templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get the requested template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = template.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.gothml from the ./templates directory
	pages, err := filepath.Glob("./templates/*.page.gothml")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.gothml
	for _, page := range pages {
		// get the file name
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
		}
		if err != nil {
			return myCache, err
		}

		myCache[name] = templateSet
	}

	return myCache, nil

}
