package handler

import "github.com/gofiber/fiber/v2"

func CheckHealth(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  "successsssss",
		"message": "Welcome to Golang, Fiber, and Postgres",
	})
}
