package repository

import (
	"errors"
	"time"

	"github.com/SamJohn04/personal-blog/src/backend/internal/config"
	"github.com/SamJohn04/personal-blog/src/backend/internal/model"
)

// Get the blog titles as a list of structs.
func GetBlogTitles() ([]model.BlogTitle, error) {
	blogTitles := []model.BlogTitle{}

	rows, err := config.DB.Query(
		"SELECT id, title, created_at, last_updated_at FROM blog",
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
		"SELECT title, html_content, created_at, last_updated_at FROM blog WHERE id=?",
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

func GetBlogToEdit(id int) (model.BlogPostEdit, error) {
	var title, content string

	row := config.DB.QueryRow(
		"SELECT title, markdown_content FROM blog WHERE id=?",
		id,
	)
	err := row.Scan(&title, &content)
	if err != nil {
		return model.BlogPostEdit{}, err
	}

	return model.BlogPostEdit{
		Id:              id,
		Title:           title,
		MarkdownContent: content,
	}, nil
}

func CreateBlogPost(title, mdContent, htmlContent string) error {
	_, err := config.DB.Exec(
		"INSERT INTO blog (title, markdown_content, html_content, created_at, last_updated_at) VALUES (?, ?, ?, ?, ?)",
		title,
		mdContent,
		htmlContent,
		time.Now(),
		time.Now(),
	)
	return err
}

func EditBlogPost(id int, title, mdContent, htmlContent string) error {
	res, err := config.DB.Exec(
		"UPDATE blog SET title = ?, markdown_content = ?, html_content = ?, last_updated_at = ? WHERE id = ?",
		title,
		mdContent,
		htmlContent,
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

func DeleteBlogPost(id int) error {
	// TODO make this function soft delete instead of hard delete
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
