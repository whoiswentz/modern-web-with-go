package config

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
)

const templatesFolder = "./templates"

var functions = template.FuncMap{}

func createTemplateCache() (map[string]*template.Template, error) {
	pagesPath := fmt.Sprintf("%s/%s", templatesFolder,"*.page.gohtml")
	basesPath := fmt.Sprintf("%s/%s", templatesFolder,"*.layout.gohtml")

	pages, err := filepath.Glob(pagesPath)
	if err != nil {
		return nil, err
	}
	log.Printf("found a total of %d pages", len(pages))

	bases, err := filepath.Glob(basesPath)
	if err != nil {
		return nil, err
	}
	log.Printf("found a total of %d bases layout", len(bases))

	templateCache := make(map[string] *template.Template)
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(basesPath)
		if err != nil {
			return nil, err
		}

		templateCache[name] = ts
	}
	return templateCache, nil
}

