package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate render templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		fmt.Println("error in createTemplateCache")
		log.Fatal(err)
	}

	// get the requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("error in tc[tmpl]")
		fmt.Println(tmpl)
		fmt.Println(tc)
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		fmt.Println("error in t.Execute")
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error in buf.WriteTo")
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.gohtml from the ./templates directory
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		fmt.Println("error in filepath.Glob")
		return myCache, err
	}

	// range through all files ending with *.page.gohtml
	for _, page := range pages {
		// get the file name
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			fmt.Println("error in template.New")
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.gohtml")
		if err != nil {
			fmt.Println("error in filepath.Glob")
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				fmt.Println("error in ts.ParseGlob")
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}
