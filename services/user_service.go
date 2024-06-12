package services

import (
	"errors"
	"forum-backend/models"
	"forum-backend/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

var secretKey = "83a20ef4d44757479d9a42cd034649f6e9227f8b466085bb49077a8ec9c4d5e4"

func CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the token

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,                           // Custom claim
		"iat":      time.Now().Unix(),                    // Issued at (current time)
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Expiration time (1 hour)
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	// Sign the token with the secret key
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.Token = tokenString

	db := utils.GetDB()
	_, err = db.Exec("INSERT INTO users (username, email, password, token ) VALUES (?,?,?,?)", user.Username, user.Email, user.Password, user.Token)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(user *models.User) (string, error) {
	db := utils.GetDB()

	var dbUser models.User

	err := db.QueryRow("SELECT id, username, email, password, token FROM users WHERE email =?", user.Email).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Email, &dbUser.Password, &dbUser.Token)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	return dbUser.Token, nil
}

func GetUsers() ([]models.User, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, username, email, password, token FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Token)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func MostPostedUser() (models.User, error) {
	db := utils.GetDB()
	var user models.User
	err := db.QueryRow("SELECT id, username, email, password, token FROM users WHERE id = (SELECT user_id FROM posts GROUP BY user_id ORDER BY COUNT(*) DESC LIMIT 1)").Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Token)
	if err != nil {
		return user, err
	}
	return user, nil
}
