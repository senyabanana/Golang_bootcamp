package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"ex0x/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

// JWTKey - секретный ключ для подписи JWT.
var JWTKey = []byte("my_secret_key")

// Credentials - структура для хранения учетных данных пользователя.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims = структура для хранения JWT claims (полезной нагрузки токена).
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetTokenHandler генерирует JWT токен при правильных учетных данных.
func GetTokenHandler(store types.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		// Простая проверка учетных данных.
		if creds.Username != "admin" || creds.Password != "password" {
			http.Error(rw, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)
		issuedTime := time.Now()
		claims := Claims{
			Username: creds.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
				IssuedAt:  jwt.NewNumericDate(issuedTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(JWTKey)
		if err != nil {
			logrus.Errorf("Error creating JWT token: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(rw).Encode(map[string]string{"token": tokenString})
	}
}
