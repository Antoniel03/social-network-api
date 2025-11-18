package handler

import (
	"encoding/json"
	"errors"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
	"net/http"
)

func (u *Handler) CreateFriendRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Create friend request handler reached")
	payload := storage.FriendRequest{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling create friend request: ", err)
		return
	}

	if err = validateFriendRequestData(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Invalid post data: ", err)
		return
	}

	ctx := r.Context()
	err = u.Service.SendFriendRequest(&payload, ctx)
	if err != nil {
		http.Error(w, "Invalid user data.", http.StatusInternalServerError)
		log.Println("Error during friend request creation: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (u *Handler) GetFriendRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Friend request handler reached")
	id := r.PathValue("id")

	//TODO make a utilities module for validation functions
	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	post, err := u.Service.GetFriendRequest(id, ctx)
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

func (u *Handler) GetUserFriendRequests(w http.ResponseWriter, r *http.Request) {
	log.Println("Get friend requests by user  handler reached")
	id := r.PathValue("id")

	//TODO make a utilities module for validation functions
	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	posts, err := u.Service.GetUserFriendRequests(id, ctx)
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

func (u *Handler) GetSentFriendRequests(w http.ResponseWriter, r *http.Request) {
	log.Println("Get friend requests by user  handler reached")
	id := r.PathValue("id")

	//TODO make a utilities module for validation functions
	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	posts, err := u.Service.GetSentFriendRequests(id, ctx)
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

func (u *Handler) DeleteFriendRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete friend request handler reached")
	id := r.PathValue("id")

	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	err := u.Service.DeleteFriendRequest(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *Handler) AcceptFriendRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Accept friend request handler reached")
	id := r.PathValue("id")

	// if isNumber(id) {
	// 	http.Error(w, "Invalid id", http.StatusBadRequest)
	// 	return
	// }

	ctx := r.Context()
	err := u.Service.AcceptFriendRequest(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func validateFriendRequestData(fr *storage.FriendRequest) error {
	if fr.SenderUserID < 1 {
		return errors.New("Invalid sender id")
	} else if fr.ReceiverUserID < 1 {
		return errors.New("Invalid receiver id")
	}
	return nil

}
