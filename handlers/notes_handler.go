package handlers

import (
	"time"

	"github.com/Fiber-CRUD/core"
	"github.com/Fiber-CRUD/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DataSource core.DataSource

func (dataSource *DataSource) GetAllNotes(c *fiber.Ctx) error {
	var id uuid.UUID
	var createdAt time.Time
	var updatedAt time.Time
	var content string
	var todos []models.Note

	rows, err := dataSource.Query(
		"SELECT * FROM notes",
	)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": err,
		})
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&id,
			&createdAt,
			&updatedAt,
			&content,
		)
		todos = append(todos, models.Note{
			Id:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Content:   content,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"notes": todos})
}

func (dataSource *DataSource) AddNote(c *fiber.Ctx) error {
	_, err := dataSource.Exec("INSERT into notes VALUES ($1, $2, $3, $4)",
		uuid.New(),
		time.Now(),
		time.Now(),
		"First Note",
	)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Note Created Successfully",
	})
}
