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

	notes, err := dataSource.QueryAllNotes()
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"notes": &notes,
	})
}

func (dataSource *DataSource) AddNote(c *fiber.Ctx) error {
	note := new(models.Note)

	if err := c.BodyParser(note); err != nil {
		c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err,
		})
	}

	fmt.Println(note.Content)

	_, err := dataSource.Exec("INSERT INTO notes VALUES ($1, $2, $3, $4)",
		uuid.New(),
		time.Now(),
		time.Now(),
		&note.Content,
	)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "Note Created Successfully",
	})
}
