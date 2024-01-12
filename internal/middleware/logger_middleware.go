package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

func LoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//DATE TIME | METHOD | PATH
		log.Infof("%s | %s | %s", time.Now(), r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
