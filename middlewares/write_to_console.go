package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("[%s] %s - %s",r.RemoteAddr, r.Method, r.URL))
		next.ServeHTTP(w, r)
	})
}
