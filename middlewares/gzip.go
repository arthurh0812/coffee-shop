package middlewares

import "net/http"

func GzipResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		next.ServeHTTP(w, req)
	})
}
