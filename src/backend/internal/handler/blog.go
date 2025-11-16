package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SamJohn04/personal-blog/src/backend/internal/middleware"
	"github.com/SamJohn04/personal-blog/src/backend/internal/repository"
	"github.com/SamJohn04/personal-blog/src/backend/internal/services"
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

	if r.URL.Query().Get("edit") != "true" {
		blog, err := repository.GetBlogPost(id)
		if err != nil {
			log.Println("Error: Database error or blog missing:", err)
			http.Error(w, "Blog is missing", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(blog)
	} else {
		blog, err := repository.GetBlogToEdit(id)
		if err != nil {
			log.Println("Error: Database error or blog missing:", err)
			http.Error(w, "Blog is missing", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(blog)
	}
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	authLevel, err := middleware.GetUserAuth(r)
	if err != nil || authLevel != 3 {
		log.Println("Error: auth level check failed. Error:", err, "; Auth level:", authLevel)
		http.Error(w, "level check failed", http.StatusUnauthorized)
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error: while decoding body:", err)
		http.Error(w, "error while decoding body", http.StatusBadRequest)
		return
	}

	htmlContent, err := services.MarkdownToHTML(req.Content)
	if err != nil {
		log.Println("Error: convertion of content failed:", err)
		http.Error(w, "conversion of body failed", http.StatusBadRequest)
		return
	}

	sanitizedHTML := services.SanitizeHTML(htmlContent)
	err = repository.CreateBlogPost(req.Title, req.Content, sanitizedHTML)
	if err != nil {
		log.Println("Error: creating blog post failed:", err)
		http.Error(w, "create blog post failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func EditBlog(w http.ResponseWriter, r *http.Request) {
	authLevel, err := middleware.GetUserAuth(r)
	if err != nil || authLevel != 3 {
		log.Println("Error: auth level check failed. Error:", err, "; Auth level:", authLevel)
		http.Error(w, "level check failed", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Error: Blog ID is missing:", err)
		http.Error(w, "Blog ID is missing", http.StatusNotFound)
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error: while decoding body:", err)
		http.Error(w, "error while decoding body", http.StatusBadRequest)
		return
	}
	err = repository.EditBlogPost(id, req.Title, req.Content)
	if err != nil {
		log.Println("Error: editing blog post failed:", err)
		http.Error(w, "edit blog post failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	authLevel, err := middleware.GetUserAuth(r)
	if err != nil || authLevel != 3 {
		log.Println("Error: auth level check failed. Error:", err, "; Auth level:", authLevel)
		http.Error(w, "level check failed", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Error: Blog ID is missing:", err)
		http.Error(w, "Blog ID is missing", http.StatusNotFound)
		return
	}
	err = repository.DeleteBlogPost(id)
	if err != nil {
		log.Println("Error: deleting blog post failed:", err)
		http.Error(w, "delete blog post failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
