package main

import (
	"fmt"
	"github.com/whoiswentz/modern-web-with-go/config"
	"github.com/whoiswentz/modern-web-with-go/handlers"
	"github.com/whoiswentz/modern-web-with-go/helpers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var appConfig config.AppConfig

	tc, err := helpers.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache", err)
	}

	appConfig.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println(fmt.Sprintf("startin app on port %s", portNumber))
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Panic(err)
	}
}

