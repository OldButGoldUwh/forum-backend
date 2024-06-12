// forum/services/comment_service.go

package services

import (
	"forum-backend/models"
	"forum-backend/utils"
)

func CreateComment(comment *models.Comment) error {
	db := utils.GetDB()
	_, err := db.Exec("INSERT INTO comments (post_id, content, user_id) VALUES (?, ?, ?)", comment.PostID, comment.Content, comment.UserID)
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
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.UserID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
