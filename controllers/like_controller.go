package controllers

import (
	"encoding/json"
	"fmt"
	"forum-backend/models"
	"forum-backend/services"
	"forum-backend/utils"
	"net/http"
)

func CreateLikeForPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CONTROLLER Create like for post")
	var like models.Like

	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	postId, postIdErr := utils.GetIdFromURL(r.RequestURI)
	if postIdErr != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	userId, _ := utils.GetUserId(token)

	like.UserID = userId
	like.PostID = postId

	err := services.CreateLikeForPost(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(like)
}

func CreateLikeForComment(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.GetUserId(token)
	like.UserID = userId

	err := services.CreateLikeForComment(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(like)
}

func CreateDislikeForPost(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.GetUserId(token)
	like.UserID = userId

	err := services.CreateDislikeForPost(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(like)
}

func CreateDislikeForComment(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.GetUserId(token)
	like.UserID = userId

	err := services.CreateDislikeForComment(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(like)
}

func UpdateLikeForPost(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.GetUserId(token)
	like.UserID = userId

	err := services.UpdateLikeForPost(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(like)
}

func UpdateLikeForComment(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.GetUserId(token)
	like.UserID = userId

	err := services.UpdateLikeForComment(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(like)
}

func DeleteLikeForPost(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.GetUserId(token)
	like.UserID = userId

	err := services.DeleteLikeForPost(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(like)
}

func DeleteLikeForComment(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.GetUserId(token)
	like.UserID = userId

	err := services.DeleteLikeForComment(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(like)
}

func GetLikes(w http.ResponseWriter, r *http.Request) {
	likes, err := services.GetLikes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likes)
}

func GetLikesForPost(w http.ResponseWriter, r *http.Request) {
	likes, err := services.GetLikesForPost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likes)
}

func GetLikesForComment(w http.ResponseWriter, r *http.Request) {
	likes, err := services.GetLikesForComment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likes)
}
