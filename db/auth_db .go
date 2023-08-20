package db

import (
	"github.com/Fiber-CRUD/types/forms"
	"github.com/Fiber-CRUD/types/models"
)

func AuthenticateUser(loginForm forms.Login) (*models.User, error) {
	user := new(models.User)

	rows, err := DBConnection.Query("SELECT * FROM users WHERE email = $1 AND password = $2")
	if err != nil {
		return user, err
	}
	defer rows.Close()

	return user, nil
}
