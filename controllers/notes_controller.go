package handler

import (
	"log"
	"time"

	"github.com/Fiber-CRUD/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type selectParams struct {
	Id uuid.UUID `json:"id" params:"id"`
}

func (dataSource *DataSource) GetAllNotes(c *fiber.Ctx) error {

	notes, err := dataSource.queryAllNotes()
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"notes": &notes,
	})
}

func (dataSource *DataSource) CreateNote(c *fiber.Ctx) error {

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

	err := dataSource.executeInsertNote(note)
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

func (dataSource *DataSource) DeleteNote(c *fiber.Ctx) error {

	id := new(selectParams)

	if err := c.ParamsParser(id); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	err := dataSource.executeDeleteNote(&id.Id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Note DELETED Successfully",
	})
}

func (dataSource *DataSource) GetNoteById(c *fiber.Ctx) error {

	id := new(selectParams)

	if err := c.ParamsParser(id); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	note, err := dataSource.queryGetNoteById(&id.Id)
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

func (dataSource *DataSource) UpdateNote(c *fiber.Ctx) error {
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

	err := dataSource.executeUpdateNote(note)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Note UPDATED Successfully",
	})
}
