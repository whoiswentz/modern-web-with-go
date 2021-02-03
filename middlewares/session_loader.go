package middlewares

import (
	"github.com/whoiswentz/modern-web-with-go/config"
	"net/http"
)

func SessionLoader(next http.Handler) http.Handler {
	return config.ApplicationConfig.Session.LoadAndSave(next)
}

