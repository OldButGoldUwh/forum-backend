// forum/cmd/main.go

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

	r.HandleFunc("/api/v1/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/api/v1/posts", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/api/v1/posts", controllers.GetPosts).Methods("GET")
	r.HandleFunc("/api/v1/posts/{id}", controllers.GetPost).Methods("GET")
	r.HandleFunc("/api/v1/posts/{title}", controllers.GetPostByTitle).Methods("GET")
	r.HandleFunc("/api/v1/posts/{id}/comments", controllers.CreateComment).Methods("POST")
	r.HandleFunc("/api/v1/most-posted-user", controllers.MostPostedUser).Methods("GET")
	r.HandleFunc("/api/v1/most-liked-posts", controllers.TenMostPopularPosts).Methods("GET")

	// Middleware
	r.Use(middlewares.AuthMiddleware)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
