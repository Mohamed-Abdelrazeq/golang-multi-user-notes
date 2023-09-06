package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/multi-user-notes-app/db"
	"github.com/multi-user-notes-app/helpers"
)

func Login(c *fiber.Ctx) error {
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

	user, err := db.DBConnection.DB.GetUserByEmail(c.Context(), authenticateUserParams.Email)
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

	tokenString, err := helpers.CreateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func Register(c *fiber.Ctx) error {

	// ALOCATE PARAMS
	createUserParams := new(db.CreateUserParams)

	// PASE PARAMS
	if err := c.BodyParser(createUserParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	// HASH PASSWORD
	password, err := helpers.HashPassword(createUserParams.Password)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	createUserParams.Password = password

	// ADD TO DB
	user, err := db.DBConnection.DB.CreateUser(c.Context(), *createUserParams)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	// WELCOME EMAIL
	go helpers.SendWelcomeEmail(createUserParams.Email)

	// SEND STATUS 200
	return c.Status(200).JSON(&fiber.Map{
		"user": user,
	})
}
