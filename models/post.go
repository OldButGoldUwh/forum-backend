// forum/models/post.go

package models

type Post struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	UserID     int      `json:"user_id"`
	Likes      int      `json:"likes"`
	Dislikes   int      `json:"dislikes"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}
