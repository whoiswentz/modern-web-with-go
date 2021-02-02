package controllers

import (
	"github.com/whoiswentz/modern-web-with-go/helpers"
	"net/http"
)

type ApplicationController struct{}

func NewApplicationController() *ApplicationController {
	return &ApplicationController{}
}

func (ac ApplicationController) Home(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "home.page.gohtml", nil)
}

func (ac ApplicationController) About(w http.ResponseWriter, r *http.Request) {
	testData := make(map[string]string)
	testData["message"] = "Hello again"

	helpers.RenderTemplate(w, "about.page.gohtml", testData)
}
