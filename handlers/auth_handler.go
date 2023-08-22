package handler

import (
	"github.com/Fiber-CRUD/db"
	"github.com/Fiber-CRUD/types/forms"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {

	loginForm := new(forms.Login)

	if err := c.BodyParser(loginForm); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := hashPassword(&loginForm.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "error hashing the password",
		})
	}

	user, err := db.GetUser(*loginForm)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Register(c *fiber.Ctx) error {

	loginForm := new(forms.Login)

	if err := c.BodyParser(loginForm); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := hashPassword(&loginForm.Password); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	err := db.CreateUser(loginForm)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "User registerd corrently",
	})
}

func Validate(c *fiber.Ctx) error {

	return nil
}

func hashPassword(password *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		return err
	}

	*password = string(bytes)
	return nil
}
