package middleware

import (
	"net/http"
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

type Decorator func(http.HandlerFunc) http.HandlerFunc

func Decorate(h http.HandlerFunc, decorators ...Decorator) http.HandlerFunc {
	for _, decorator := range decorators {
		h = decorator(h)
	}
	return h
}
