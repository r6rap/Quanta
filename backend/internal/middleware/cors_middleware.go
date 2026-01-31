package middleware

import (
	"net/http"

	"github.com/r6rap/Quanta/internal/handler"
)

type CORSMiddleware struct {
	handler *handler.AuthHandler
}

func NewCORSMiddleware(handler *handler.AuthHandler) *CORSMiddleware {
	return &CORSMiddleware{handler: handler}
}

func (c *CORSMiddleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}