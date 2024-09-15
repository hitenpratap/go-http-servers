package middleware

import (
	"fmt"
	"net/http"
)

// LoggingMiddleware logs the requested URL and passes control to the next handler
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requested URL: %s\n", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
