// forum/services/comment_service.go

package services

import (
	"forum-backend/models"
	"forum-backend/utils"
)

func CreateComment(comment *models.Comment) error {
	db := utils.GetDB()
	_, err := db.Exec("INSERT INTO comments (post_id, content, user_id,created_at) VALUES (?, ?, ?,?)", comment.PostID, comment.Content, comment.UserID, comment.CreatedAt)
	return err
}

func GetComments() ([]models.Comment, error) {
	db := utils.GetDB()

	rows, err := db.Query("SELECT * FROM comments")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.UserID, &comment.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func GetCommentsForPost(postId int) ([]models.Comment, error) {
	db := utils.GetDB()

	rows, err := db.Query("SELECT * FROM comments WHERE post_id =?", postId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.UserID, &comment.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func GetComment(id int) (*models.Comment, error) {
	db := utils.GetDB()
	var comment models.Comment
	err := db.QueryRow("SELECT * FROM comments WHERE id = ?", id).Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.UserID)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func DeleteComment(id int) error {
	db := utils.GetDB()
	_, err := db.Exec("DELETE FROM comments WHERE id =?", id)
	return err
}

func UpdateComment(comment *models.Comment) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE comments SET content =? WHERE id =?", comment.Content, comment.ID)
	return err
}

func MostCommentedPost() (int, error) {
	db := utils.GetDB()
	var postId int
	err := db.QueryRow("SELECT post_id FROM comments GROUP BY post_id ORDER BY COUNT(post_id) DESC LIMIT 1").Scan(&postId)
	if err != nil {
		return 0, err
	}
	return postId, nil
}
