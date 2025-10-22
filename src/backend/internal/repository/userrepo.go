package repository

import (
	"github.com/SamJohn04/personal-blog/src/backend/internal/config"
	"github.com/SamJohn04/personal-blog/src/backend/internal/model"
)

func CreateUser(user model.User) error {
	_, err := config.DB.Exec(
		"INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)",
		user.Username, user.Email, user.Password,
	)
	return err
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	user.Email = email

	err := config.DB.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE email=?",
		email,
	).Scan(&user.Id, &user.Username, &user.Password)

	return user, err
}
