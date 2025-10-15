package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SamJohn04/personal-blog/src/backend/internal/repository"
)

func GetBlogTitles(w http.ResponseWriter, r *http.Request) {
	blogTitles := repository.GetBlogTitles()
	json.NewEncoder(w).Encode(blogTitles)
}
