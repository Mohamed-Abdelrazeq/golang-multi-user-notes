package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/multi-user-notes-app/db/internals"
	"github.com/multi-user-notes-app/db/models"
	"github.com/multi-user-notes-app/helpers"
)

func Login(c *fiber.Ctx) error {
	loginParams := new(models.LoginParams)

	if err := c.BodyParser(loginParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := helpers.Validator.Struct(loginParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid email or password",
		})
	}

	user, err := internals.DBConnection.DB.GetUserByEmail(c.Context(), loginParams.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Invalid email",
		})
	}

	if isValid := helpers.CheckPasswordHash(loginParams.Password, user.Password); !isValid {
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
	registerParams := new(models.RegisterParams)

	// PASE PARAMS
	if err := c.BodyParser(registerParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := helpers.Validator.Struct(registerParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// HASH PASSWORD
	password, err := helpers.HashPassword(registerParams.Password)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	registerParams.Password = password

	// ADD TO DB
	dbRegisterParams := internals.CreateUserParams{
		Email:    registerParams.Email,
		Password: registerParams.Password,
	}
	user, err := internals.DBConnection.DB.CreateUser(c.Context(), dbRegisterParams)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	// WELCOME EMAIL
	go helpers.SendWelcomeEmail(registerParams.Email)

	// SEND STATUS 200
	return c.Status(200).JSON(&fiber.Map{
		"user": user,
	})
}

// TODO: HARD CODED FOR NOW
func DeleteUserById(c *fiber.Ctx) error {
	err := internals.DBConnection.DB.DeleteUserById(c.Context(), 2)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "User deleted successfully",
	})
}
