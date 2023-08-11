package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Route("/api", func(router fiber.Router) {
		router.Get("/notes", getAllNotes)
		router.Post("/notes", addNote)
		router.Get("/health", checkHealth)
	})

	app.Listen("127.0.0.1:8080")
}
