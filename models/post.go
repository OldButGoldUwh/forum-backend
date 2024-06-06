// forum/models/post.go

package models

type Post struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	UserID     int      `json:"user_id"`
}
