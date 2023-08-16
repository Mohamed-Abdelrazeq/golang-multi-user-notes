package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (dataSource *DataSource) GetAllNotes(c *fiber.Ctx) error {

	notes, err := dataSource.QueryAllNotes()
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"notes": &notes,
	})
}

// func (dataSource *DataSource) AddNote(c *fiber.Ctx) error {
// 	note := new(models.Note)

// 	if err := c.BodyParser(note); err != nil {
// 		c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
// 			"message": err,
// 		})
// 	}

// 	log.Fatalln(note)

// 	_, err := dataSource.Exec("INSERT into notes VALUES ($1, $2, $3, $4)",
// 		uuid.New(),
// 		time.Now(),
// 		time.Now(),
// 		"First Note",
// 	)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
// 			"message": err,
// 		})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"message": "Note Created Successfully",
// 	})
// }
