package handler

import (
	"time"

	"github.com/Fiber-CRUD/db"
	"github.com/Fiber-CRUD/types/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUser(c *fiber.Ctx) error {

	note := models.Note{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// TODO: FIND A BETTER WAY TO VALIDATE MODELS
	if err := c.BodyParser(&note); err != nil || note.Content == "" {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": "Content can't be empty",
		})
	}

	err := db.InsertNote(note)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "Note Created Successfully",
		"note":    note,
	})
}
