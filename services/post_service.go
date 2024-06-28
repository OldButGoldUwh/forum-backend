// forum/services/post_service.go

package services

import (
	"fmt"
	"forum-backend/models"
	"forum-backend/utils"
)

func CreatePost(post *models.Post, userId int) error {

	db := utils.GetDB()
	post.CreatedAt = utils.GetCurrentTime()
	categories := ""

	for i, category := range post.Categories {
		if i == 0 {
			categories = category
		} else {
			categories = categories + "," + category
		}
	}

	_, err := db.Exec("INSERT INTO posts (title, content, user_id, categories,likes, dislikes, comment_length, created_at, updated_at ) VALUES (?, ?, ?, ?,?,?,?,?,?)", post.Title, post.Content, post.UserID, categories, 0, 0, 0, post.CreatedAt, post.CreatedAt)

	return err
}

func GetPosts() ([]models.Post, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, title, content, user_id, comment_length, likes, dislikes, updated_at, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CommentLength, &post.Likes, &post.Dislikes, &post.UpdatedAt, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func EditPost(post *models.Post) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE posts SET title =?, content =?, updated_at =? WHERE id =?", post.Title, post.Content, utils.GetCurrentTime(), post.ID)
	return err
}

func UpdatePostCommentLength(postId int) error {
	db := utils.GetDB()
	var CommentLength int
	errS := db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id =?", postId).Scan(&CommentLength)
	if errS != nil {
		return errS
	}
	_, err := db.Exec("UPDATE posts SET  updated_at =?, comment_length =? WHERE id =?", utils.GetCurrentTime(), CommentLength+1, postId)
	return err
}

func GetPost(id int) (*models.Post, error) {
	db := utils.GetDB()
	var post models.Post
	err := db.QueryRow(("Select id, title, content, user_id, created_at, updated_at FROM posts WHERE id =?"), id).Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
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

func UpdatePostUpdatedAt(postId int) error {
	db := utils.GetDB()
	currentTime := utils.GetCurrentTime()
	_, err := db.Exec("UPDATE posts SET updated_at =? WHERE id =?", currentTime, postId)
	fmt.Println("Updated at: ", currentTime)
	return err

}

func UpdatePostLikes(postId int) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE posts SET likes = likes + 1 WHERE id =?", postId)
	return err
}

func UpdatePostDislikes(postId int) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE posts SET dislikes = dislikes + 1 WHERE id =?", postId)
	return err
}

func DeletePostLikes(postId int) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE posts SET likes = likes - 1 WHERE id =?", postId)
	return err
}

func DeletePostDislikes(postId int) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE posts SET dislikes = dislikes - 1 WHERE id =?", postId)
	return err
}

func GetPostUserId(postId int) (int, error) {
	db := utils.GetDB()
	var userId int
	err := db.QueryRow("SELECT user_id FROM posts WHERE id =?", postId).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func DeletePost(postId int) error {
	db := utils.GetDB()
	_, err := db.Exec("DELETE FROM posts WHERE id =?", postId)
	return err
}

func GetPostComments(postId int) ([]models.Comment, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT id, content, user_id FROM comments WHERE post_id =?", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
