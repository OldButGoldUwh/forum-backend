// forum/models/comment.go

package models

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}
