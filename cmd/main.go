package main

import (
	"log"
	"net/http"

	"forum-backend/controllers"
	"forum-backend/middlewares"
	"forum-backend/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.InitDB()
	r := mux.NewRouter()
	// Auth
	r.HandleFunc("/api/v1/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", controllers.UpdateUser).Methods("UPDATE")
	r.HandleFunc("/api/v1/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// User
	r.HandleFunc("/api/v1/users/{id}/posts", controllers.GetUserPosts).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}/comments", controllers.GetUserComments).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}/likes", controllers.GetUserLikes).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}/dislikes", controllers.GetUserDislikes).Methods("GET")
	r.HandleFunc("/api/v1/most-posted-user", controllers.MostPostedUser).Methods("GET")
	r.HandleFunc("/api/v1/profile-data", controllers.ProfileData).Methods("GET")

	// Comment
	r.HandleFunc("/api/v1/posts/{id}/comments", controllers.CreateComment).Methods("POST")
	r.HandleFunc("/api/v1/posts/{id}/comments", controllers.GetComments).Methods("GET")
	r.HandleFunc("/api/v1/comments/{id}", controllers.GetComment).Methods("GET")
	r.HandleFunc("/api/v1/comments/{id}", controllers.UpdateComment).Methods("UPDATE")
	r.HandleFunc("/api/v1/comments/{id}", controllers.DeleteComment).Methods("DELETE")

	// Posts
	r.HandleFunc("/api/v1/posts", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/api/v1/posts", controllers.GetPosts).Methods("GET")
	r.HandleFunc("/api/v1/posts/{id}", controllers.GetPost).Methods("GET")
	r.HandleFunc("/api/v1/posts/{title}", controllers.GetPostByTitle).Methods("GET")
	r.HandleFunc("/api/v1/most-liked-posts", controllers.TenMostPopularPosts).Methods("GET")
	r.HandleFunc("/api/v1/posts/{id}", controllers.EditPost).Methods("UPDATE")
	r.HandleFunc("/api/v1/posts/{id}", controllers.DeletePost).Methods("DELETE")

	// Like and dislike
	r.HandleFunc("/api/v1/posts/{id}", controllers.CreateLikeForComment).Methods("POST")
	r.HandleFunc("/api/v1/posts/{id}", controllers.CreateLikeForPost).Methods("POST")
	r.HandleFunc("/api/v1/posts/{id}", controllers.CreateDislikeForComment).Methods("POST")
	r.HandleFunc("/api/v1/posts/{id}", controllers.CreateDislikeForPost).Methods("POST")
	r.HandleFunc("/api/v1/posts/{id}", controllers.DeleteLikeForComment).Methods("DELETE")
	r.HandleFunc("/api/v1/posts/{id}", controllers.DeleteLikeForPost).Methods("DELETE")
	r.HandleFunc("/api/v1/posts/{id}", controllers.UpdateLikeForComment).Methods("UPDATE")
	r.HandleFunc("/api/v1/posts/{id}", controllers.UpdateLikeForPost).Methods("UPDATE")
	r.HandleFunc("/api/v1/likes", controllers.GetLikes).Methods("GET")
	r.HandleFunc("/api/v1/likes", controllers.GetLikesForPost).Methods("GET")
	r.HandleFunc("/api/v1/likes", controllers.GetLikesForComment).Methods("GET")

	// Search
	r.HandleFunc("/api/v1/search/users/{username}", controllers.SearchUsers).Methods("GET")
	r.HandleFunc("/api/v1/search/posts/{title}", controllers.SearchPosts).Methods("GET")
	r.HandleFunc("/api/v1/search/comments/{content}", controllers.SearchComments).Methods("GET")

	// Middleware
	r.Use(middlewares.AuthMiddleware)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
