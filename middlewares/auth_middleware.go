package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey []byte = []byte("gokhan")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware çalıştı")
		// Authorization başlığını al
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Yetkilendirme hatası: Token eksik", http.StatusUnauthorized)
			return
		}

		// "Bearer <token>" formatını kontrol et
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(w, "Yetkilendirme hatası: Token formatı hatalı", http.StatusUnauthorized)
			return
		}

		tokenStr := splitToken[1]

		// Token doğrulama
		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Yetkilendirme hatası: İmza geçersiz", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Yetkilendirme hatası: "+err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Yetkilendirme hatası: Token geçersiz", http.StatusUnauthorized)
			return
		}

		// Token geçerliyse bir sonraki handler'a geç
		next.ServeHTTP(w, r)
	})
}
