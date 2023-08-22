package db

import (
	"errors"
	"time"

	"github.com/Fiber-CRUD/types/forms"
	"github.com/Fiber-CRUD/types/models"
)

func GetUser(loginForm forms.Login) (*models.User, error) {

	user := new(models.User)
	rows, err := DBConnection.Query("SELECT * FROM users WHERE email = $1", loginForm.Email)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("user doesn't exist")
	}

	rows.Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Password,
	)

	return user, nil
}

func CreateUser(loginForm *forms.Login) error {
	_, err := DBConnection.Exec(
		"INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		loginForm.Email,
		loginForm.Password,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
