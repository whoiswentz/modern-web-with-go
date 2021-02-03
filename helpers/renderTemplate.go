package helpers

import (
	"bytes"
	"github.com/whoiswentz/modern-web-with-go/config"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, file string, data interface{}) {
	tmpl, ok := config.ApplicationConfig.Templates[file]
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
