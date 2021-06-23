package middlewares

import (
	"log"
	"net/http"
)

func Print(l *log.Logger, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		l.Printf("%s %s", req.Method, req.URL.Path)
		handler.ServeHTTP(w, req)
	})
}
