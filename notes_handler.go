package main

import "github.com/gofiber/fiber/v2"

type Note struct {
	Content string `json:"content"`
}

func getAllNotes(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"notes": ""})
}

func addNote(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Note Created Successfully"})
}
