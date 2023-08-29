package supa

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nedpals/supabase-go"
)

type user struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SUPABASE IS SO SLOW 999

func CreateUser(c *fiber.Ctx) error {

	type params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	row := params{
		Email:    "mohamed21@gmail.com",
		Password: "1234567",
	}

	var results []user

	go SupabaseClient.Auth.SignUp(ctx, supabase.UserCredentials{
		Email:    row.Email,
		Password: row.Password,
	})

	SupabaseClient.DB.From("users").Insert(row).Execute(&results)

	return nil
}
