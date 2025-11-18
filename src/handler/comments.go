package handler

import (
	"encoding/json"
	"errors"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
	"net/http"
	// "strconv"
	// "unicode"
)

func (u *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	log.Println("Create comment handler reached")
	payload := storage.Comment{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling create comment: ", err)
		return
	}

	if err = validateCommentData(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Invalid user data: ", err)
		return
	}

	ctx := r.Context()
	err = u.Service.PublishComment(&payload, ctx)
	if err != nil {
		http.Error(w, "Invalid comment data.", http.StatusInternalServerError)
		log.Println("Error during comment publication: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (u *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	log.Println("Get comment handler reached")
	id := r.PathValue("id")

	if isNumber(id) {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	user, err := u.Service.GetComment(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *Handler) GetPostComments(w http.ResponseWriter, r *http.Request) {
	log.Println("Get post comments handler reached")
	id := r.PathValue("id")

	//TODO make a utilities module for validation functions
	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	posts, err := u.Service.GetCommentsByPost(id, ctx)
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

func (u *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete comment handler reached")
	id := r.PathValue("id")

	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	err := u.Service.DeleteComment(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func validateCommentData(comment *storage.Comment) error {
	if comment.TextContent == "" {
		return errors.New("Text content was not provided")
	} else if comment.UserId < 1 {
		return errors.New("Incorrect user id")
	} else if comment.PostId < 1 {
		return errors.New("Incorrect post id")
	}
	return nil
}
