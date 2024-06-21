package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	secretKey  = "83a20ef4d44757479d9a42cd034649f6e9227f8b466085bb49077a8ec9c4d5e4"
	guestToken = "0fc237962e95129004c313015d220aef4c7ffddc465cf984d1e63130b6e180c8"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Check if the token is the guest token
		if tokenString == guestToken {
			next.ServeHTTP(w, r)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			fmt.Println("Error parsing token:", err)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			fmt.Println("Token is not valid")
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func GuestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != guestToken {
			fmt.Println(authHeader)
			http.Error(w, "Invalid guest token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
