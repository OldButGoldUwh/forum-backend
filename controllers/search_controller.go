package controllers

import (
	"encoding/json"
	"forum-backend/services"
	"net/http"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	users, err := services.SearchUsers(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func SearchPosts(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	posts, err := services.SearchPosts(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func SearchComments(w http.ResponseWriter, r *http.Request) {
	content := r.URL.Query().Get("content")
	comments, err := services.SearchComments(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)

}
