package main

import (
	"log"

	"github.com/Fiber-CRUD/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Route("/api", func(router fiber.Router) {
		router.Get("/notes", controllers.GetAllNotes)
		// router.Post("/notes", controllers.AddNote)
		router.Get("/health", controllers.CheckHealth)
	})
	log.Fatal(
		app.Listen("127.0.0.1:8080"),
	)
}
