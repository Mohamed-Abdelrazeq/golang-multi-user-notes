package main

import (
	"database/sql"
	"log"

	handler "github.com/Fiber-CRUD/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	dataSource, err := openDB()
	if err != nil {
		log.Fatal("ERROR CONNECTING TO DB")
	}

	app.Route("/api", func(router fiber.Router) {
		router.Get("/notes/:id", dataSource.GetNoteById)
		router.Delete("/notes/:id", dataSource.DeleteNote)
		router.Post("/notes", dataSource.CreateNote)
		router.Patch("/notes", dataSource.UpdateNote)
		router.Get("/notes", dataSource.GetAllNotes)
		router.Get("/health", handler.CheckHealth)
	})
	log.Fatal(
		app.Listen("127.0.0.1:8080"),
	)
}

func openDB() (*handler.DataSource, error) {
	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:5024@localhost:5432/Fiber-CRUD?sslmode=disable",
	)
	if err != nil {
		return nil, err
	}

	return &handler.DataSource{DB: conn}, nil
}
