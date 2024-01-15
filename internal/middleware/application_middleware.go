package middleware

import (
	"net/http"

	"github.com/labstack/gommon/log"
)

func ApplicationMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		log.Infof("%s | %s", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
