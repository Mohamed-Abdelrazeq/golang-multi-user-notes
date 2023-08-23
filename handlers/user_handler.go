package handler

import (
	"github.com/Fiber-CRUD/db"
	"github.com/Fiber-CRUD/helpers"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateUser(c *fiber.Ctx) error {
	type AuthenticateUserParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	authenticateUserParams := new(AuthenticateUserParams)

	if err := c.BodyParser(authenticateUserParams); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := db.DBConnection.DB.AuthenticateUser(c.Context(), authenticateUserParams.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Invalid email",
		})
	}

	if isValid := helpers.CheckPasswordHash(authenticateUserParams.Password, user.Password); !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Invalid password",
		})
	}

	tokenString, err := helpers.CreateToken(int(user.ID))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func CreateUser(c *fiber.Ctx) error {

	createUserParams := new(db.CreateUserParams)

	if err := c.BodyParser(createUserParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	password, err := helpers.HashPassword(createUserParams.Password)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	createUserParams.Password = password

	user, err := db.DBConnection.DB.CreateUser(c.Context(), *createUserParams)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"user": user,
	})
}
