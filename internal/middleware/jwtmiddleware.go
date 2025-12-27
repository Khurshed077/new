package middleware

import (
	"context"
	"net/http"
	"strings"

	"newww/internal/middleware/jwt"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenString string

		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {

			cookie, err := r.Cookie("access_token")
			if err == nil && cookie.Value != "" {
				tokenString = cookie.Value
			}
		}

		if tokenString == "" {
			next.ServeHTTP(w, r)
			return
		}

		claims, err := jwt.ParseAccessToken(tokenString)
		if err != nil {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), jwt.ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
