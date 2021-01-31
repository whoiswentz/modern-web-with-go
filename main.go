package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println(fmt.Sprintf("startin app on port %s", portNumber))
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, file string) {
	filePath := fmt.Sprintf("./templates/%s", file)
	parsedTemplate, err := template.ParseFiles(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	if err := parsedTemplate.Execute(w, nil); err != nil {
		log.Println("error parsing template: ", err)
	}
}
