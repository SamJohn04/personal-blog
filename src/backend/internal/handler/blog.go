package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SamJohn04/personal-blog/src/backend/internal/middleware"
	"github.com/SamJohn04/personal-blog/src/backend/internal/repository"
	"github.com/go-chi/chi/v5"
)

func GetBlogTitles(w http.ResponseWriter, r *http.Request) {
	blogTitles, err := repository.GetBlogTitles()
	if err != nil {
		log.Println("Error: DB error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(blogTitles)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Error: Blog ID is missing:", err)
		http.Error(w, "Blog ID is missing", http.StatusNotFound)
		return
	}
	blog, err := repository.GetBlogPost(id)
	if err != nil {
		log.Println("Error: Database error or blog missing:", err)
		http.Error(w, "Blog is missing", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(blog)
}

type CreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error: while decoding body:", err)
		http.Error(w, "error while decoding body", http.StatusBadRequest)
		return
	}

	authLevel, err := middleware.GetUserAuth(r)
	if err != nil || authLevel != 3 {
		log.Println("Error: auth level check failed. Error:", err, "; Auth level:", authLevel)
		http.Error(w, "level check failed", http.StatusUnauthorized)
		return
	}
	err = repository.CreatePost(req.Title, req.Content)
	if err != nil {
		log.Println("Error: creating post failed:", err)
		http.Error(w, "create post failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

type EditRequest struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func EditBlog(w http.ResponseWriter, r *http.Request) {
	authLevel, err := middleware.GetUserAuth(r)
	if err != nil || authLevel != 3 {
		log.Println("Error: auth level check failed. Error:", err, "; Auth level:", authLevel)
		http.Error(w, "level check failed", http.StatusUnauthorized)
		return
	}

	var req EditRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error: while decoding body:", err)
		http.Error(w, "error while decoding body", http.StatusBadRequest)
		return
	}
	err = repository.EditPost(req.Id, req.Title, req.Content)
	if err != nil {
		log.Println("Error: editing post failed:", err)
		http.Error(w, "edit post failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
