package app

import (
	"log"
	"net/http"

	"github.com/SamJohn04/personal-blog/src/backend/internal/handler"
	"github.com/SamJohn04/personal-blog/src/backend/internal/middleware"
	"github.com/go-chi/chi/v5"
)

// Run links all the handlers via chi Router and serves the production database as a FileServer.
func Run() {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Logger)

		r.Post("/register", handler.RegisterUser)
		r.Post("/login", handler.LoginUser)

		r.Get("/blogs", handler.GetBlogTitles)
		r.Get("/blog/{id}", handler.GetBlog)

		r.Route("/blog", func(r chi.Router) {
			r.Use(middleware.Auth)

			r.Post("/", handler.CreateBlog)
			r.Put("/{id}", handler.EditBlog)
			r.Delete("/{id}", handler.DeleteBlog)
		})
	})

	fs := http.FileServer(http.Dir("../frontend/dist"))
	r.Handle("/*", fs)

	log.Println("Starting server on", 8000)
	http.ListenAndServe(":8000", r)
}
