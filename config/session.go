package config

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"time"
)

func createSession() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = false // For study purpose only

	return session
}