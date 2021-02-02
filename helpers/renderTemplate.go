package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const templatesFolder = "./templates"

var functions = template.FuncMap{}
var templates map[string]*template.Template

func init() {
	tmpls, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	templates = tmpls
}

func RenderTemplate(w http.ResponseWriter, file string, data interface{}) {
	tmpl, ok := templates[file]
	if !ok {
		log.Fatalln("template not found")
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		log.Fatalln(err)
	}

	if _, err := buf.WriteTo(w); err != nil {
		log.Println("error while writing template to browser", err)
	}
}

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
