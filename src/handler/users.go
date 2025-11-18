package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"unicode"

	"github.com/Antoniel03/social-network-api/internal/storage"
)

func jwtHandler(id int, name string, lastname string) (*string, error) {
	claims := CustomClaims{
		Name:     name,
		Lastname: lastname,
		Id:       id,
	}
	token, err := GenerateJWT("guacamole", claims)
	if err != nil {
		log.Printf("Error while creating jwt: %s", err)
		return nil, err
	}
	return token, nil
}

func (u *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Create user handler reached")
	payload := storage.User{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling creater user: ", err)
		return
	}

	if err = validateRegistrationData(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Invalid user data: ", err)
		return
	}

	ctx := r.Context()
	err = u.Service.RegisterUser(&payload, ctx)
	if err != nil {
		http.Error(w, "Invalid user data.", http.StatusInternalServerError)
		log.Println("Error during user registration: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (u *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Get user handler reached")
	id := r.PathValue("id")

	if isNumber(id) {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	user, err := u.Service.GetUser(id, ctx)
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
}

func (u *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Get user handler reached")
	id := r.PathValue("id")

	if isNumber(id) {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err := u.Service.DeleteUser(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *Handler) UpdateName(w http.ResponseWriter, r *http.Request) {
	log.Println("Update name handler reached")
	var payload map[string]string
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling update name: ", err)
		return
	}
	log.Println(payload)

	id := r.PathValue("id")

	if isNumber(id) {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		log.Println("Invalid id")
		return
	}

	if payload["name"] == "" || !isAlpha(payload["name"]) {
		http.Error(w, "Invalid name data.", http.StatusBadRequest)
		log.Println("Invalid name entry")
		return
	}

	ctx := r.Context()
	err = u.Service.UpdateUser("name", payload["name"], id, ctx)
	if err != nil {
		http.Error(w, "Invalid name data.", http.StatusInternalServerError)
		log.Println("Error during name update: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *Handler) UpdateLastname(w http.ResponseWriter, r *http.Request) {
	log.Println("Update lastname handler reached")
	var payload map[string]string
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling update lastname: ", err)
		return
	}
	log.Println(payload)

	id := r.PathValue("id")

	if isNumber(id) {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		log.Println("Invalid id")
		return
	}

	if payload["lastname"] == "" || !isAlpha(payload["lastname"]) {
		http.Error(w, "Invalid lastname data.", http.StatusBadRequest)
		log.Println("Invalid lastname entry")
		return
	}

	ctx := r.Context()
	err = u.Service.UpdateUser("lastname", payload["lastname"], id, ctx)
	if err != nil {
		http.Error(w, "Invalid lastname data.", http.StatusInternalServerError)
		log.Println("Error during lastname update: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *Handler) UpdateEmail(w http.ResponseWriter, r *http.Request) {
	log.Println("Update email handler reached")
	var payload map[string]string
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling update email: ", err)
		return
	}
	log.Println(payload)

	id := r.PathValue("id")

	if isNumber(id) {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	//TODO email validation

	ctx := r.Context()
	err = u.Service.UpdateUser("email", payload["email"], id, ctx)
	if err != nil {
		http.Error(w, "Invalid user data.", http.StatusInternalServerError)
		log.Println("Error during email update: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	log.Println("Update password handler reached")
	var payload map[string]string
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling update password: ", err)
		return
	}
	log.Println(payload)

	id := r.PathValue("id")

	if isNumber(id) {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if payload["password"] == "" {
		http.Error(w, "Invalid password data.", http.StatusBadRequest)
	}

	ctx := r.Context()
	err = u.Service.UpdateUser("password", payload["password"], id, ctx)
	if err != nil {
		http.Error(w, "Invalid password data.", http.StatusInternalServerError)
		log.Println("Error during password update: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isNumber(number string) bool {
	if _, err := strconv.Atoi(number); err == nil {
		return false
	}
	return true
}

func (u *Handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login handler reached")
	payload := storage.User{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error handling login: ", err)
		return
	}

	ctx := r.Context()
	user, err := u.Service.Login(payload.Email, payload.Password, ctx)
	if err != nil {
		log.Println("Error handling login: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := jwtHandler(user.Id, user.Name, user.Lastname)
	if err != nil {
		log.Println("Error handling login: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(map[string]string{"token": *token})
	if err != nil {
		log.Println("Error handling login: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func validateRegistrationData(user *storage.User) error {
	if user.Name == "" {
		return errors.New("Name was not provided")
	} else if user.Lastname == "" {
		return errors.New("Lastname was not provided")
	} else if user.BirthDate == "" {
		return errors.New("Birth date was not provided")
	} else if user.Email == "" {
		return errors.New("Email was not provided")
	} else if user.Password == "" {
		return errors.New("Password was not provided")
	}
	return nil
}
