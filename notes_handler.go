package main

import "github.com/gofiber/fiber/v2"

func getAllNotes(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"notes": []Note{
		{
			Content: "Note #1",
		},
		{
			Content: "Note #2",
		},
		{
			Content: "Note #3",
		},
	}})
}

func addNote(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Note Created Successfully"})
}
