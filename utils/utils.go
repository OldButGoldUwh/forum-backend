// forum/utils/utils.go

package utils

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var JwtKey = []byte("your_secret_key") // Replace with your actual secret key

func GetToken(r *http.Request) string {
	// Get the token from the request header
	token := r.Header.Get("Authorization")
	if token == "" {
		return ""
	}

	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		return ""
	}

	return token
}

func GetUserId(token string) (int, error) {
	db := GetDB()
	if token == "" {
		return 0, errors.New("token is empty")
	}

	claims := &jwt.StandardClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, errors.New("invalid token signature")
		}
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	// Query the database to find the user associated with the token
	var userId int
	err = db.QueryRow("SELECT id FROM users WHERE token = ?", token).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	return userId, nil
}
