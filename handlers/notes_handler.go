package handler

import (
	"github.com/Fiber-CRUD/db"
	"github.com/Fiber-CRUD/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetAllNotes(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)

	notes, err := db.DBConnection.DB.GetAllNotes(c.Context(), userId)
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
	userId := helpers.RecoverToken(c)

	if err := c.BodyParser(&createNoteParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	createNoteParams.UserID = userId

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

func DeleteNote(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)
	deleteNoteParams := new(db.DeleteNoteParams)

	if err := c.ParamsParser(deleteNoteParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	deleteNoteParams.UserID = userId

	err := db.DBConnection.DB.DeleteNote(c.Context(), *deleteNoteParams)
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

	userId := helpers.RecoverToken(c)
	getNoteByIdParams := new(db.GetNoteByIdParams)

	if err := c.ParamsParser(getNoteByIdParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	getNoteByIdParams.UserID = userId

	note, err := db.DBConnection.DB.GetNoteById(c.Context(), *getNoteByIdParams)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"note": &note,
	})
}

func UpdateNote(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)
	updateNoteParams := new(db.UpdateNoteParams)

	if err := c.BodyParser(updateNoteParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	updateNoteParams.UserID = userId

	note, err := db.DBConnection.DB.UpdateNote(c.Context(), *updateNoteParams)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"note": note,
	})
}
