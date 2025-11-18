package handler

import (
	"encoding/json"
	"errors"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
	"net/http"
)

func (u *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Create post handler reached")
	payload := storage.Post{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling create post: ", err)
		return
	}

	if err = validatePostData(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Invalid post data: ", err)
		return
	}

	ctx := r.Context()
	err = u.Service.PublishPost(&payload, ctx)
	if err != nil {
		http.Error(w, "Invalid user data.", http.StatusInternalServerError)
		log.Println("Error during post creation: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (u *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Get post handler reached")
	id := r.PathValue("id")

	//TODO make a utilities module for validation functions
	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	post, err := u.Service.GetPost(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (u *Handler) GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Get posts by user  handler reached")
	id := r.PathValue("id")

	//TODO make a utilities module for validation functions
	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	posts, err := u.Service.GetPostsDataByUser(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO fix this thing because it is allowing null results and sending 200 status
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (u *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete post handler reached")
	id := r.PathValue("id")

	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	err := u.Service.DeletePost(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func validatePostData(post *storage.Post) error {
	if post.UserId == 0 {
		return errors.New("User id was not provided")
	} else if post.TextContent == "" {
		return errors.New("Text content was not provided")
	}
	return nil
}

func (u *Handler) LikePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Like post handler reached")
	userId := r.PathValue("userid")
	postId := r.PathValue("postid")
	log.Printf("%s - %s", userId, postId)

	//TODO make a utilities module for validation functions
	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	err := u.Service.AddLike(userId, postId, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *Handler) DislikePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Dislike post handler reached")
	userId := r.PathValue("userid")
	postId := r.PathValue("postid")

	log.Printf("userid: %s, postid: %s", userId, postId)

	//TODO make a utilities module for validation functions
	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	err := u.Service.AddDislike(userId, postId, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
