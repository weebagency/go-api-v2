package middleware

import (
	"log"
	"net/http"
)

func WithLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logger middleware running\n")
		next.ServeHTTP(w, r)
	}
}
