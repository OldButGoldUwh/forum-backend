package services

import (
	"forum-backend/models"
	"forum-backend/utils"
)

func SearchUsers(username string) ([]models.User, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, username, email, password, token FROM users WHERE username LIKE ?", username)
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

func SearchPosts(title string) ([]models.Post, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, title, content, user_id FROM posts WHERE title LIKE ?", title)
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

func SearchComments(content string) ([]models.Comment, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, post_id, content, user_id FROM comments WHERE content LIKE ?", content)
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
