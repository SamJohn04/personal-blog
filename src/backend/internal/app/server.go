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

	r.Post("/register", handler.RegisterUser)
	r.Post("/login", handler.Login)

	r.Get("/blogs", handler.GetBlogTitles)
	r.Get("/blog/{id}", handler.GetBlog)

	r.Route("/blog", func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Post("/", handler.CreateBlog)
		r.Put("/{id}", handler.EditBlog)
	})

	log.Println("Starting server on", 8000)
	http.ListenAndServe(":8000", r)
}
