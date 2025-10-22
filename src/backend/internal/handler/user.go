package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SamJohn04/personal-blog/src/backend/internal/model"
	"github.com/SamJohn04/personal-blog/src/backend/internal/repository"
	"github.com/SamJohn04/personal-blog/src/backend/internal/utils"
)

type AuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashed,
	}
	err = repository.CreateUser(user)
	if err != nil {
		http.Error(w, "username and/or email already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "user does not exist", http.StatusBadRequest)
		return
	} else if req.Username != user.Username {
		http.Error(w, "wrong username", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
