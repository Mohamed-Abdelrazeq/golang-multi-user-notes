package handler

import (
	"log"
	"time"

	"github.com/Fiber-CRUD/db"
	"github.com/Fiber-CRUD/types/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type selectParams struct {
	Id uuid.UUID `json:"id" params:"id"`
}

func GetAllNotes(c *fiber.Ctx) error {

	notes, err := db.GetAllNotes()
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"notes": &notes,
	})
}

func CreateNote(c *fiber.Ctx) error {

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

	err := db.CreateNote(note)
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

func DeleteNote(c *fiber.Ctx) error {

	id := new(selectParams)

	if err := c.ParamsParser(id); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	err := db.DeleteNote(&id.Id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Note DELETED Successfully",
	})
}

func GetNoteById(c *fiber.Ctx) error {

	id := new(selectParams)

	if err := c.ParamsParser(id); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	note, err := db.GetNoteById(&id.Id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "note is found",
		"note":    &note,
	})
}

func UpdateNote(c *fiber.Ctx) error {
	// TODO: PARSE ID FROM PARAMS
	note := models.Note{
		UpdatedAt: time.Now(),
	}

	// TODO: FIND A BETTER WAY TO VALIDATE MODELS
	if err := c.BodyParser(&note); err != nil || note.Content == "" {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": "Content can't be empty",
		})
	}

	err := db.UpdateNote(note)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Note UPDATED Successfully",
	})
}
