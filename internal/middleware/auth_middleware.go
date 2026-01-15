package middleware

import (
	"context"
	"net/http"

	"github.com/r6rap/Quanta/internal/handler"
	"github.com/r6rap/Quanta/pkg/jwt"
	"github.com/r6rap/Quanta/pkg/response"
)

type AuthMiddleware struct {
	handler *handler.AuthHandler
}

func NewAuthMiddleware(handler *handler.AuthHandler) *AuthMiddleware {
	return &AuthMiddleware{handler: handler}
}

func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := jwt.ExtractToken(r)

		if token == "" {
			response.Fail(w, http.StatusUnauthorized, "UNAUTHORIZED", "Missing authorization header")
			return
		}

		claims, err := jwt.ValidateToken(token)
		if err != nil {
			response.Fail(w, http.StatusInternalServerError, "INTERNAL_ERROR", "error validate token")
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}