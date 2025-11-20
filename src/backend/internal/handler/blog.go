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

// GetBlogTitles returns the blog titles as an encoded list of structs. May return an error code
func GetBlogTitles(w http.ResponseWriter, r *http.Request) {
	blogTitles, err := repository.GetBlogTitles()
	if err != nil {
		log.Println("Error: DB error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(blogTitles)
}

// GetBlog returns a single blog as a struct. May return an error code.
// Use with query edit = true to get the markdown content.
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

// Create Blog returns a status code if it is successful.
// Can return an error code.
// Auth level 3 (Author) required.
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

// Edit Blog returns with a status code.
// Can return an error code.
// Auth level 3 (Author) required.
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

	htmlContent, err := services.MarkdownToHTML(req.Content)
	if err != nil {
		log.Println("Error: convertion of content failed:", err)
		http.Error(w, "conversion of body failed", http.StatusBadRequest)
		return
	}

	sanitizedHTML := services.SanitizeHTML(htmlContent)
	err = repository.EditBlogPost(id, req.Title, req.Content, sanitizedHTML)
	if err != nil {
		log.Println("Error: editing blog post failed:", err)
		http.Error(w, "edit blog post failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Blog returns with a status code.
// Can return an error code.
// Auth level 3 (Author) required.
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
