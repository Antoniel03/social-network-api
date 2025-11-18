package main

import (
	"log"
	"net/http"

	"github.com/Antoniel03/social-network-api/handler"
	"github.com/Antoniel03/social-network-api/internal/storage"
)

type Logger struct {
	handler http.Handler
}

type Authenticator struct {
	handler func(w http.ResponseWriter, r *http.Request)
}

func verifyAuth(handler func(w http.ResponseWriter, r *http.Request)) Authenticator {
	return Authenticator{handler}
}

func (a Authenticator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Protected Route, verifying credentials")
	authArgument := r.Header.Get("Authorization")

	if err := handler.ValidateAuthorizedRequest(authArgument); err != nil {
		log.Println(err)
	} else {
		a.handler(w, r)
	}
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	l.handler.ServeHTTP(w, r)
}

type Service struct {
	xd storage.Storage
}

func SetupRouter(router *http.ServeMux, a *handler.Handler) {
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Helloooooo, i hope you're doing great!\n"))

	})
	router.HandleFunc("POST /login", a.Login)
	router.Handle("POST /users", verifyAuth(a.CreateUser))
	router.Handle("GET /users/{id}", verifyAuth(a.GetUser))
	router.Handle("DELETE /users/{id}", verifyAuth(a.Delete))
	router.Handle("PATCH /users/{id}/name", verifyAuth(a.UpdateName))
	router.Handle("PATCH /users/{id}/lastname", verifyAuth(a.UpdateLastname))
	router.Handle("PATCH /users/{id}/email", verifyAuth(a.UpdateEmail))
	router.Handle("PATCH /users/{id}/password", verifyAuth(a.UpdatePassword))

	router.Handle("POST /posts", verifyAuth(a.CreatePost))
	router.Handle("GET /posts/{id}", verifyAuth(a.GetPost))
	router.Handle("GET /posts/users/{id}", verifyAuth(a.GetPostsByUser))
	router.Handle("POST /posts/{postid}/users/{userid}/like", verifyAuth(a.LikePost))
	router.Handle("POST /posts/{postid}/users/{userid}/dislike", verifyAuth(a.DislikePost))
	router.Handle("DELETE /posts/{id}", verifyAuth(a.DeletePost))

	//TODO
	// router.HandleFunc("POST /media", a.CreateMedia)

	router.Handle("POST /comments", verifyAuth(a.CreateComment))
	router.Handle("GET /comments/{id}", verifyAuth(a.GetComment))
	router.Handle("GET /comments/posts/{id}", verifyAuth(a.GetPostComments))
	router.Handle("DELETE /comments/{id}", verifyAuth(a.DeleteComment))

	router.Handle("POST /friend-requests", verifyAuth(a.CreateFriendRequest))
	router.Handle("PATCH /friend-requests/{id}", verifyAuth(a.AcceptFriendRequest))
	router.Handle("GET /friend-requests/{id}", verifyAuth(a.GetFriendRequest))
	router.Handle("GET /friend-requests/users/{id}", verifyAuth(a.GetUserFriendRequests))
	router.Handle("GET /friend-requests/sender/{id}", verifyAuth(a.GetSentFriendRequests))
	router.Handle("DELETE /friend-requests/{id}", verifyAuth(a.DeleteFriendRequest))
}
