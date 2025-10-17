package repository

import (
	"time"

	"github.com/SamJohn04/personal-blog/src/backend/internal/config"
	"github.com/SamJohn04/personal-blog/src/backend/internal/model"
)

func GetBlogTitles() ([]model.BlogTitle, error) {
	blogTitles := []model.BlogTitle{}

	rows, err := config.DB.Query(
		"SELECT id, title, createdAt, lastUpdatedAt FROM blog",
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
	var title, content string
	var createdAt, lastUpdatedAt time.Time

	row := config.DB.QueryRow(
		"SELECT title, content, createdAt, lastUpdatedAt FROM blog WHERE id=?",
		id,
	)
	err := row.Scan(&title, &content, &createdAt, &lastUpdatedAt)
	if err != nil {
		return model.BlogPost{}, err
	}

	return model.BlogPost{
		Id:            id,
		Title:         title,
		Content:       content,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
	}, nil
}
