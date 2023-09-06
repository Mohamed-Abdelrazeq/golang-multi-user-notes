package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/multi-user-notes-app/db"
)

// TODO: HARD CODED FOR NOW
func DeleteUserById(c *fiber.Ctx) error {
	err := db.DBConnection.DB.DeleteUserById(c.Context(), 1)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "User deleted successfully",
	})
}
