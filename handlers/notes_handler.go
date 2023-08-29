package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/multi-user-notes-app/connections"
	"github.com/multi-user-notes-app/db"
	"github.com/multi-user-notes-app/helpers"
)

func GetAllNotes(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)

	notes, err := connections.DBConnection.DB.GetAllNotes(c.Context(), userId)
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

	params := new(db.CreateNoteParams)
	userId := helpers.RecoverToken(c)

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	note, err := connections.DBConnection.DB.CreateNote(c.Context(), *params)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": params,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"note": note,
	})
}

func DeleteNote(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)
	params := new(db.DeleteNoteParams)

	if err := c.ParamsParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	err := connections.DBConnection.DB.DeleteNote(c.Context(), *params)
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
	params := new(db.GetNoteByIdParams)

	if err := c.ParamsParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	note, err := connections.DBConnection.DB.GetNoteById(c.Context(), *params)
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
	params := new(db.UpdateNoteParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	note, err := connections.DBConnection.DB.UpdateNote(c.Context(), *params)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"note": note,
	})
}

func AddToFavourites(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)
	params := new(db.AddToFavouritesParams)

	if err := c.ParamsParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	note, err := connections.DBConnection.DB.AddToFavourites(c.Context(), *params)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"note": note,
	})
}

func RemoveFromFavourite(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)
	params := new(db.RemoveFromFavouritesParams)

	if err := c.ParamsParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	note, err := connections.DBConnection.DB.RemoveFromFavourites(c.Context(), *params)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"note": note,
	})
}

func ListFavourites(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)

	notes, err := connections.DBConnection.DB.ListFavourites(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"notes": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"notes": notes,
	})
}
