package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/whoiswentz/modern-web-with-go/controllers"
	"github.com/whoiswentz/modern-web-with-go/middlewares"
	"log"
	"net/http"
	"time"
)

const (
	portNumber = ":8080"
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 90 * time.Second
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
)

func main() {
	appController := controllers.NewApplicationController()

	router := mux.NewRouter()
	router.Use(middlewares.WriteToConsole)

	router.HandleFunc("/", appController.Home)
	router.HandleFunc("/about", appController.About)

	srv := &http.Server{
		Addr:              portNumber,
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	log.Println(fmt.Sprintf("startint app on port %s", portNumber))
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

