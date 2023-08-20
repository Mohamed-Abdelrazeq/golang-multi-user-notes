package handler

import (
	"time"

	"github.com/Fiber-CRUD/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	type LoginForm struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	loginForm := new(LoginForm)

	if err := c.BodyParser(&loginForm); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	password, err := hashPassword(loginForm.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "error hashing the password",
		})
	}

	loginForm.Password = password

	rows, err := db.DBConnection.Query("SELECT * FROM users WHERE email = $1 AND password = $2")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "user is not found",
		})
	}

	for rows.Next() {

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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
