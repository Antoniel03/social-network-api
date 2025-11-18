package handler

import (
	"encoding/json"
	// "errors"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"log"
	"net/http"
	// "strconv"
	// "unicode"
)

func (u *Handler) CreateMedia(w http.ResponseWriter, r *http.Request) {
	log.Println("Create media handler reached")
	payload := storage.Media{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling create media: ", err)
		return
	}

	// if err = validateMediaData(&payload); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	log.Println("Invalid media data: ", err)
	// 	return
	// }

	ctx := r.Context()
	err = u.Service.SaveMedia(&payload, ctx)
	if err != nil {
		http.Error(w, "Invalid media data.", http.StatusInternalServerError)
		log.Println("Error during media save: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
