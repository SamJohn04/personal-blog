package repository

import (
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

func GetBlogPost(id int) model.BlogPost {
	return blogPosts[id]
}
