package controllers

import (
	"forum-backend/models"
	"forum-backend/services"
)

func SearchUsers(username string) ([]models.User, error) {
	users, err := services.SearchUsers(username)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func SearchPosts(title string) ([]models.Post, error) {
	posts, err := services.SearchPosts(title)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func SearchComments(content string) ([]models.Comment, error) {
	comments, err := services.SearchComments(content)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
