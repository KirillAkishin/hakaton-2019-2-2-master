package main

import (
	"client/pkg"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"math/rand"
	"net/http"
	"rclone2/pkg/handlers"
	"rclone2/pkg/middleware"
	"rclone2/pkg/post"
	"rclone2/pkg/session"
	"rclone2/pkg/user"
	"time"
)

func main() {
	fmt.Println()
	rand.Seed(time.Now().UnixNano())
	templates := template.Must(template.ParseFiles("./template/index.html"))
	staticHandler := http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./template/static/")),
	)
	http.Handle("/static/", staticHandler)

	userRepoz := user.NewUserRepo()
	postRepoz := post.NewPostRepo()
	sesManager := session.NewSesManager()
	userHandler := &pkg.UserHandler{
		Tmpl:     templates,
		UserRepo: &userRepoz,
		Sessions: &sesManager,
	}
	postHandler := &handlers.PostHandler{
		PostRepo: &postRepoz,
	}

	r := mux.NewRouter()

	r.HandleFunc("/", userHandler.Template).Methods("GET")
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/api/register", userHandler.Register).Methods("POST")

	r.HandleFunc("/api/posts", postHandler.NewPost).Methods("POST")
	r.HandleFunc("/api/post/{id}", postHandler.ShowPost).Methods("GET")
	r.HandleFunc("/api/posts/", postHandler.ShowAllPosts).Methods("GET")
	r.HandleFunc("/api/posts/{category}", postHandler.ShowAllCategoryPosts).Methods("GET")
	r.HandleFunc("/api/user/{user_login}", postHandler.ShowAllUserPosts).Methods("GET")
	r.HandleFunc("/api/post/{id}", postHandler.DeletePost).Methods("DELETE")

	r.HandleFunc("/api/post/{post_id}", postHandler.NewComment).Methods("POST")
	r.HandleFunc("/api/post/{post_id}/{comment_id}", postHandler.DeleteComment).Methods("DELETE")

	r.HandleFunc("/api/post/{post_id}/upvote", postHandler.Upvote).Methods("GET")
	r.HandleFunc("/api/post/{post_id}/downvote", postHandler.Downvote).Methods("GET")
	r.HandleFunc("/api/post/{post_id}/unvote", postHandler.Unvote).Methods("GET")

	muxH := middleware.Auth(&sesManager, r)
	muxH = middleware.Panic(muxH)
	http.ListenAndServe(":8080", muxH)
}
