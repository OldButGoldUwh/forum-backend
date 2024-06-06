// forum/services/user_service.go

package services

import (
	"errors"
	"forum-backend/models"
	"forum-backend/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("your_secret_key")

func CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	db := utils.GetDB()
	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?,?,?)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(user *models.User) (string, error) {
	db := utils.GetDB()

	var dbUser models.User

	err := db.QueryRow("SELECT id, username, email, password FROM users WHERE email =?", user.Email).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Email, &dbUser.Password)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil {
		return "", errors.New("KULLANICI ADI VEYA SIFRE YANLISTIR")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = strconv.Itoa(dbUser.ID)
	claims["username"] = dbUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func GetUsers() ([]models.User, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}
