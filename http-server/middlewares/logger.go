package middlewares

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(logger *log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%s - %s - %s", r.Host, r.RequestURI, r.Method)
			next.ServeHTTP(w, r)
		})
	}
}
