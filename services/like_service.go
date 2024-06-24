package services

import (
	"fmt"
	"forum-backend/models"
	"forum-backend/utils"
)

func CreateLikeForPost(like *models.Like) error {
	fmt.Println("SERVİCE Create like for post")
	db := utils.GetDB()
	_, err := db.Exec("INSERT INTO likes (post_id, user_id, created_at, like) VALUES (?, ?, ?, ?)", like.PostID, like.UserID, like.CreatedAt, true)
	return err
}

func CreateLikeForComment(like *models.Like) error {
	db := utils.GetDB()
	_, err := db.Exec("INSERT INTO likes (comment_id, user_id, created_at, like) VALUES (?,?,?,?)", like.CommentID, like.UserID, like.CreatedAt, true)
	return err
}

func CreateDislikeForPost(like *models.Like) error {
	db := utils.GetDB()
	_, err := db.Exec("INSERT INTO likes (post_id, user_id, created_at, like) VALUES (?, ?, ?, ?)", like.PostID, like.UserID, like.CreatedAt, false)
	return err
}

func CreateDislikeForComment(like *models.Like) error {
	db := utils.GetDB()
	_, err := db.Exec("INSERT INTO likes (comment_id, user_id, created_at, like) VALUES (?,?,?,?)", like.CommentID, like.UserID, like.CreatedAt, false)
	return err
}

func DeleteLikeForPost(like *models.Like) error {
	db := utils.GetDB()
	_, err := db.Exec("DELETE FROM likes WHERE post_id = ? AND user_id = ?", like.PostID, like.UserID)
	return err
}

func DeleteLikeForComment(like *models.Like) error {
	db := utils.GetDB()
	_, err := db.Exec("DELETE FROM likes WHERE comment_id = ? AND user_id = ?", like.CommentID, like.UserID)
	return err
}

func UpdateLikeForPost(like *models.Like) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE likes SET like = ? WHERE post_id = ? AND user_id = ?", like.Like, like.PostID, like.UserID)
	return err
}

func UpdateLikeForComment(like *models.Like) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE likes SET like = ? WHERE comment_id = ? AND user_id = ?", like.Like, like.CommentID, like.UserID)
	return err
}

func GetLikes() ([]models.Like, error) {
	db := utils.GetDB()

	rows, err := db.Query("SELECT * FROM likes")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []models.Like

	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.ID, &like.PostID, &like.CommentID, &like.UserID, &like.CreatedAt, &like.Like)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	return likes, nil
}

func GetLikesForPost() ([]models.Like, error) {
	db := utils.GetDB()

	rows, err := db.Query("SELECT * FROM likes WHERE like = true")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []models.Like

	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.ID, &like.PostID, &like.CommentID, &like.UserID, &like.CreatedAt, &like.Like)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	return likes, nil
}

func GetLikesForComment() ([]models.Like, error) {
	db := utils.GetDB()

	rows, err := db.Query("SELECT * FROM likes WHERE like = true")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []models.Like

	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.ID, &like.PostID, &like.CommentID, &like.UserID, &like.CreatedAt, &like.Like)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	return likes, nil
}
