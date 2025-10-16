package app

import (
	"log"
	"net/http"

	"github.com/SamJohn04/personal-blog/src/backend/internal/handler"
	"github.com/SamJohn04/personal-blog/src/backend/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func Run() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/blogs", handler.GetBlogTitles)
	r.Get("/blog/{id}", handler.GetBlog)

	log.Println("Starting server on", 8000)
	http.ListenAndServe(":8000", r)
}
