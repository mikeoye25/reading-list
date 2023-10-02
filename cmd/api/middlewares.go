package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusForbidden)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "Malformed token", http.StatusForbidden)
			return
		}

		tokenPart := parts[1]
		token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("supersecretkey"), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			http.Error(w, "Token is not valid", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
