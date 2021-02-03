package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

var ApplicationConfig *appConfig

func init() {
	ApplicationConfig = newAppConfig()
}

type appConfig struct {
	Templates map[string]*template.Template
	Session *scs.SessionManager
}

func newAppConfig() *appConfig {
	tmpls, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	return &appConfig{
		Templates: tmpls,
		Session: createSession(),
	}
}