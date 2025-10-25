package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
