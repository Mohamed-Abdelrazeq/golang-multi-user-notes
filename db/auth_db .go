package db

import (
	"time"

	"github.com/Fiber-CRUD/types/forms"
)

// func AuthenticateUser(loginForm forms.Login) (*models.User, error) {
// 	user := new(models.User)

// 	rows, err := DBConnection.Query("SELECT * FROM users WHERE email = $1 AND password = $2")
// 	if err != nil {
// 		return user, err
// 	}
// 	defer rows.Close()

// 	return user, nil
// }

func CreateUser(loginForm *forms.Login) error {
	_, err := DBConnection.Exec(
		"INSERT INTO users (email, password, created_at, updated_at, token) VALUES ($1, $2, $3, $4, $5)",
		loginForm.Email,
		loginForm.Password,
		time.Now(),
		time.Now(),
		"",
	)

	if err != nil {
		return err
	}

	return nil
}
