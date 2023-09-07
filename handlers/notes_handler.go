package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/multi-user-notes-app/db/internals"
	"github.com/multi-user-notes-app/db/models"
	"github.com/multi-user-notes-app/helpers"
)

func GetAllNotes(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)

	notes, err := internals.DBConnection.DB.GetAllNotes(c.Context(), userId)
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

	params := new(internals.CreateNoteParams)
	userId := helpers.RecoverToken(c)

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := helpers.Validator.Struct(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	dbParams := internals.CreateNoteParams{
		UserID:      userId,
		Title:       params.Title,
		Content:     params.Content,
		ImageUrl:    params.ImageUrl,
		IsFavourite: params.IsFavourite,
	}

	note, err := internals.DBConnection.DB.CreateNote(c.Context(), dbParams)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"note": note,
	})
}

func DeleteNote(c *fiber.Ctx) error {

	userId := helpers.RecoverToken(c)
	params := new(models.NoteDetailsParams)

	if err := c.ParamsParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := helpers.Validator.Struct(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	dbParams := internals.DeleteNoteParams{
		ID:     params.ID,
		UserID: userId,
	}

	err := internals.DBConnection.DB.DeleteNote(c.Context(), dbParams)
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
	params := new(models.NoteDetailsParams)

	if err := c.ParamsParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := helpers.Validator.Struct(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	dbParams := internals.GetNoteByIdParams{
		ID:     params.ID,
		UserID: userId,
	}

	note, err := internals.DBConnection.DB.GetNoteById(c.Context(), dbParams)
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
	params := new(internals.UpdateNoteParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	note, err := internals.DBConnection.DB.UpdateNote(c.Context(), *params)
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
	params := new(internals.AddToFavouritesParams)

	if err := c.ParamsParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	note, err := internals.DBConnection.DB.AddToFavourites(c.Context(), *params)
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
	params := new(internals.RemoveFromFavouritesParams)

	if err := c.ParamsParser(params); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	params.UserID = userId

	note, err := internals.DBConnection.DB.RemoveFromFavourites(c.Context(), *params)
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

	notes, err := internals.DBConnection.DB.ListFavourites(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"notes": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"notes": notes,
	})
}
