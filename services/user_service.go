package services

import (
	"errors"
	"forum-backend/models"
	"forum-backend/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

// TODO SESSİON KULLANILACAK

var secretKey = "83a20ef4d44757479d9a42cd034649f6e9227f8b466085bb49077a8ec9c4d5e4"

func CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	time := time.Now().Unix()

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"expaires": time + 3600,
		// TODO SESSİON KULLANILACAK
		"iat": time,
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

func GetUser(id int) (models.User, error) {
	db := utils.GetDB()
	var user models.User
	err := db.QueryRow("SELECT id, username, email, password, token FROM users WHERE id =?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Token)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	db := utils.GetDB()
	_, err = db.Exec("UPDATE users SET username =?, email =?, password =? WHERE id =?", user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	db := utils.GetDB()
	_, err := db.Exec("DELETE FROM users WHERE id =?", id)
	return err
}

func GetUserPosts(id int) ([]models.Post, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, title, content, user_id FROM posts WHERE user_id =?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetUserComments(id int) ([]models.Comment, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, post_id, content, user_id FROM comments WHERE user_id =?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.UserID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func GetUserLikes(id int) ([]models.Like, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, post_id, comment_id, user_id, created_at, like FROM likes WHERE user_id =?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []models.Like
	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.ID, &like.PostID, &like.CommentID, &like.UserID, &like.CreatedAt, &like.Like)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return likes, nil
}

func GetUserDislikes(id int) ([]models.Like, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, post_id, comment_id, user_id, created_at, like FROM likes WHERE user_id =? AND like = 0", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []models.Like
	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.ID, &like.PostID, &like.CommentID, &like.UserID, &like.CreatedAt, &like.Like)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return likes, nil
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
