package config

import (
	"html/template"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

