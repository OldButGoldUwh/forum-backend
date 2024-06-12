// forum/services/post_service.go

package services

import (
	"forum-backend/models"
	"forum-backend/utils"
)

func CreatePost(post *models.Post) error {
	// authMiddleware'de token kontrolü yapılacak

	db := utils.GetDB()

	_, err := db.Exec("INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)", post.Title, post.Content)
	return err
}

func GetPosts() ([]models.Post, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, title, content, user_id FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPost(id int) (*models.Post, error) {
	db := utils.GetDB()
	var post models.Post
	err := db.QueryRow("SELECT id, title, content, user_id FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPostByTitle(title string) (*models.Post, error) {
	db := utils.GetDB()
	var post models.Post
	err := db.QueryRow("SELECT id, title, content, user_id FROM posts WHERE title =?", title).Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func TenMostPopularPosts() ([]models.Post, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, title, content, user_id FROM posts ORDER BY likes DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func MostLikedPost() (models.Post, error) {
	db := utils.GetDB()
	var post models.Post
	err := db.QueryRow("SELECT id, title, content, user_id FROM posts ORDER BY likes DESC LIMIT 1").Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func RecentTopic() (models.Post, error) {
	db := utils.GetDB()
	var post models.Post
	err := db.QueryRow("SELECT id, title, content, user_id FROM posts ORDER BY created_at DESC LIMIT 1").Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}
