package repository

import (
	"errors"
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

func CreateBlogPost(title, content string) error {
	_, err := config.DB.Exec(
		"INSERT INTO blog (title, content, createdAt, lastUpdatedAt) VALUES (?, ?, ?, ?)",
		title,
		content,
		time.Now(),
		time.Now(),
	)
	return err
}

func EditBlogPost(id int, title, content string) error {
	res, err := config.DB.Exec(
		"UPDATE blog SET title = ?, content = ?, lastUpdatedAt = ? WHERE id = ?",
		title,
		content,
		time.Now(),
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("blog not found")
	}
	return nil
}

// TODO make this function soft delete instead of hard delete
func DeleteBlogPost(id int) error {
	res, err := config.DB.Exec("DELETE FROM blog WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("blog not found")
	}
	return nil
}
