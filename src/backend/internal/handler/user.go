package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SamJohn04/personal-blog/src/backend/internal/model"
	"github.com/SamJohn04/personal-blog/src/backend/internal/repository"
	"github.com/SamJohn04/personal-blog/src/backend/internal/utils"
)

type RegisterAuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterAuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Println("Error: Could not hash password:", err)
		http.Error(w, "could not hash password", http.StatusInternalServerError)
		return
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashed,
	}
	err = repository.CreateUser(user)
	if err != nil {
		log.Println("Error: Registeration failed:", err)
		http.Error(w, "username and/or email already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginAuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("Error: User does not exist:", err)
		http.Error(w, "user does not exist", http.StatusBadRequest)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		log.Println("Error: Passwords do not match, or something went wrong")
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
