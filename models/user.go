// forum/models/user.go

package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
