package repository

import (
	"errors"
	"time"

	"github.com/SamJohn04/personal-blog/src/backend/internal/model"
)

// TODO convert to db
var blogPosts = []model.BlogPost{
	{
		Id:            1,
		Title:         "Hello World",
		Content:       "Hello, World!\nHow are you?\n",
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
	},
	{
		Id:            2,
		Title:         "Title 2",
		Content:       "Hi there ^.^",
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
	},
}

func GetBlogTitles() []model.BlogTitle {
	blogTitles := []model.BlogTitle{}
	for _, blogPost := range blogPosts {
		blogTitles = append(blogTitles, model.BlogTitle{
			Id:            blogPost.Id,
			Title:         blogPost.Title,
			CreatedAt:     blogPost.CreatedAt,
			LastUpdatedAt: blogPost.LastUpdatedAt,
		})
	}
	return blogTitles
}

func GetBlogPost(id int) (model.BlogPost, error) {
	for _, blogPost := range blogPosts {
		if blogPost.Id == id {
			return blogPost, nil
		}
	}
	return model.BlogPost{}, errors.New("not found")
}
