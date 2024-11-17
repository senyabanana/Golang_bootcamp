package middleware

import (
	"context"
	"net/http"
	"strings"

	"ex0x/internal/handlers"
	"github.com/golang-jwt/jwt/v5"
)

// Тип для ключей контекста, чтобы избежать конфликтов.
type contextKey string

const contextKeyUsername = contextKey("username")

// JWTMiddleware проверяет валидность JWT-токена и добавляет информацию о пользователе в контекст запроса.
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(rw, "Missing Authorization header", http.StatusUnauthorized)
			return
		}
		// Извлечение токена из заголовка Authorization.
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &handlers.Claims{}

		// Парсинг токена и проверка его подписи.
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return handlers.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(rw, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			http.Error(rw, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(rw, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Добавление имени пользователя в контекст запроса.
		ctx := context.WithValue(r.Context(), contextKeyUsername, claims.Username)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
