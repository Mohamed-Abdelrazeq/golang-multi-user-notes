package handler

import (
	"github.com/Fiber-CRUD/db"
	"github.com/Fiber-CRUD/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetAllNotes(c *fiber.Ctx) error {

	id := helpers.RecoverToken(c)

	notes, err := db.DBConnection.DB.GetAllNotes(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"notes": &notes,
	})
}

func CreateNote(c *fiber.Ctx) error {

	createNoteParams := new(db.CreateNoteParams)
	id := helpers.RecoverToken(c)

	if err := c.BodyParser(&createNoteParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	createNoteParams.UserID = id

	note, err := db.DBConnection.DB.CreateNote(c.Context(), *createNoteParams)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": createNoteParams,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"note": note,
	})
}

// func DeleteNote(c *fiber.Ctx) error {

// 	id := new(selectParams)

// 	if err := c.ParamsParser(id); err != nil {
// 		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	err := db.DeleteNote(&id.Id)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 		"message": "Note DELETED Successfully",
// 	})
// }

// func GetNoteById(c *fiber.Ctx) error {

// 	id := new(selectParams)

// 	if err := c.ParamsParser(id); err != nil {
// 		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	note, err := db.GetNoteById(&id.Id)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 		"message": "note is found",
// 		"note":    &note,
// 	})
// }

// func UpdateNote(c *fiber.Ctx) error {
// 	// TODO: PARSE ID FROM PARAMS
// 	note := models.Note{
// 		UpdatedAt: time.Now(),
// 	}

// 	// TODO: FIND A BETTER WAY TO VALIDATE MODELS
// 	if err := c.BodyParser(&note); err != nil || note.Content == "" {
// 		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
// 			"message": "Content can't be empty",
// 		})
// 	}

// 	err := db.UpdateNote(note)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 		"message": "Note UPDATED Successfully",
// 	})
// }
