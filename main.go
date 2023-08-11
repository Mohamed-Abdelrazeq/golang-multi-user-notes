package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	database, err := sql.Open("postgres", "postgres://postgres:5024@localhost:5432/Fiber-CRUD?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	dataSource := DataSource{database}

	app.Route("/api", func(router fiber.Router) {
		router.Get("/notes", dataSource.getAllNotes)
		router.Post("/notes", dataSource.addNote)
		router.Get("/health", checkHealth)
	})

	log.Fatal(app.Listen("127.0.0.1:8080"))
}
