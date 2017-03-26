package logging

import (
	"log"
	"net/http"
	"time"
)

func Middleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			duration := time.Since(start)
			log.Printf("%s %s %s\n", r.Method, r.URL.Path, duration.String())
		})
	}
}
