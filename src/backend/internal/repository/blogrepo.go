package repository

import (
	"errors"
	"time"

	"github.com/SamJohn04/personal-blog/src/backend/internal/config"
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

func GetBlogTitles() ([]model.BlogTitle, error) {
	blogTitles := []model.BlogTitle{}

	rows, err := config.DB.Query(
		"SELECT id, title, createdAt, LastUpdatedAt FROM blog",
	)
	if err != nil {
		return blogTitles, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var createdAt, lastUpdatedAt time.Time

		rows.Scan(&id, &title, &createdAt, &lastUpdatedAt)
		blogTitles = append(blogTitles, model.BlogTitle{
			Id:            id,
			Title:         title,
			CreatedAt:     createdAt,
			LastUpdatedAt: lastUpdatedAt,
		})
	}

	return blogTitles, nil
}

func GetBlogPost(id int) (model.BlogPost, error) {
	for _, blogPost := range blogPosts {
		if blogPost.Id == id {
			return blogPost, nil
		}
	}
	return model.BlogPost{}, errors.New("not found")
}
