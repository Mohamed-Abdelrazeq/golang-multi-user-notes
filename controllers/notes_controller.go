package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/Fiber-CRUD/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (dataSource *DataSource) GetAllNotes(c *fiber.Ctx) error {

	notes, err := dataSource.queryAllNotes()
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"notes": &notes,
	})
}

func (dataSource *DataSource) AddNote(c *fiber.Ctx) error {

	note := models.Note{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// TODO: FIND A BETTER WAY TO MARSHAL JSON
	if err := c.BodyParser(&note); err != nil {
		c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err,
		})
	}

	fmt.Println(len(note.Content))

	err := dataSource.excuteInsertNote(note)
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
