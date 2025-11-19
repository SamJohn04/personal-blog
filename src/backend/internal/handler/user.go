package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SamJohn04/personal-blog/src/backend/internal/model"
	"github.com/SamJohn04/personal-blog/src/backend/internal/repository"
	"github.com/SamJohn04/personal-blog/src/backend/internal/utils"
)

// RegisterUser checks the username, email, and password, hashes the password, and stores them.
// The function will return either an error code or an empty body.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error: while decoding body:", err)
		http.Error(w, "error while decoding body", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		log.Println("Error: The username, email or password is empty")
		http.Error(w, "empty username or email or password", http.StatusBadRequest)
		return
	}

	hashed, err := utils.GenerateHashFromPassword(req.Password)
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

// LoginUser logs in an existing user.
// The function returns either an error code or the string encoding of the token and auth level (an integer from 0 to 3).
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error: while decoding body:", err)
		http.Error(w, "error while decoding body", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("Error: User does not exist:", err)
		http.Error(w, "user does not exist", http.StatusBadRequest)
		return
	}

	if !utils.CheckPasswordWithHash(req.Password, user.Password) {
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
		"token":     token,
		"authLevel": strconv.Itoa(user.AuthLevel),
	})
}
