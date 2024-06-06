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
