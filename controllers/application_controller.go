package controllers

import (
	"github.com/whoiswentz/modern-web-with-go/config"
	"github.com/whoiswentz/modern-web-with-go/helpers"
	"net/http"
)

type ApplicationController struct{}

func NewApplicationController() *ApplicationController {
	return &ApplicationController{}
}

func (ac ApplicationController) Home(w http.ResponseWriter, r *http.Request) {
	config.GetAppConfig().Session.Put(r.Context(), "remote_ip", r.RemoteAddr)

	helpers.RenderTemplate(w, "home.page.gohtml", nil)
}

func (ac ApplicationController) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := config.GetAppConfig().Session.GetString(r.Context(), "remote_ip")

	testData := make(map[string]string)
	testData["remote_ip"] = remoteIP
	testData["message"] = "Hello"

	helpers.RenderTemplate(w, "about.page.gohtml", testData)
}
