// forum/utils/utils.go

package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var JwtKey = []byte("83a20ef4d44757479d9a42cd034649f6e9227f8b466085bb49077a8ec9c4d5e4") // Replace with your actual secret key

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

	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		return 0, errors.New("token is empty")
	}

	claims := &jwt.StandardClaims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("invalid token signature")
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

func GetCurrentTime() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func GetIdFromURL(url string) (int, error) {
	parts := strings.Split(url, "/")
	for _, part := range parts {
		if id, err := strconv.Atoi(part); err == nil {
			return id, nil
		}
	}
	return 0, errors.New("ID not found in URL")
}
