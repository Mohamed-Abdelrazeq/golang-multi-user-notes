package handler

import (
	"time"

	"github.com/Fiber-CRUD/db"
	"github.com/Fiber-CRUD/types/forms"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(c *fiber.Ctx) error {

	loginForm := new(forms.Login)

	if err := c.BodyParser(&loginForm); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := hashPassword(&loginForm.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "error hashing the password",
		})
	}

	_, err := db.AuthenticateUser(*loginForm)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "user is not found",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func hashPassword(password *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		return err
	}

	*password = string(bytes)
	return nil
}
