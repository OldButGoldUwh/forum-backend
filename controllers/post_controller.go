// forum/controllers/post_controller.go

package controllers

import (
	"encoding/json"
	"forum-backend/models"
	"forum-backend/services"
	"forum-backend/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.GetUserId(token)

	err := services.CreatePost(&post, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := services.GetPost(id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func GetPostByTitle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	title := params["title"]

	post, err := services.GetPostByTitle(title)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func TenMostPopularPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := services.TenMostPopularPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func MostLikedPost(w http.ResponseWriter, r *http.Request) {
	post, err := services.MostLikedPost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func RecentTopic(w http.ResponseWriter, r *http.Request) {
	post, err := services.RecentTopic()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func UpdatePostLikes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"]) // Post ID
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	userId, _ := utils.GetUserId(r.Header.Get("Authorization"))
	if userId == 0 {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = services.UpdatePostLikes(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, _ := utils.GetUserId(r.Header.Get("Authorization"))

	if userId == 0 {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	getUserIdFromPost, err := services.GetPostUserId(id)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if userId != getUserIdFromPost {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = services.EditPost(&post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	userId, _ := utils.GetUserId(r.Header.Get("Authorization"))

	if userId == 0 {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	getUserIdFromPost, err := services.GetPostUserId(id)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if userId != getUserIdFromPost {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = services.DeletePost(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func MostCommentedPost(w http.ResponseWriter, r *http.Request) {
	post, err := services.MostCommentedPost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
