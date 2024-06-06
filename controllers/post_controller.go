// forum/controllers/post_controller.go

package controllers

import (
	"encoding/json"
	"forum-backend/models"
	"forum-backend/services"
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

	err := services.CreatePost(&post)
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
