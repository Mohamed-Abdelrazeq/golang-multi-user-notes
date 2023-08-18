package main

import (
	"log"

	handler "github.com/Fiber-CRUD/controllers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	// MIDDLEWARES
	app.Use(logger.New())
	app.Use("/api", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))
	// DB INIT
	dataSource, err := handler.OpenDB()
	if err != nil {
		log.Fatal("ERROR CONNECTING TO DB")
	}

	app.Route("/authenticate", func(router fiber.Router) {
		router.Post("/login", dataSource.Login)
	})

	app.Route("/api", func(router fiber.Router) {
		// NOTES
		router.Get("/notes/:id", dataSource.GetNoteById)
		router.Delete("/notes/:id", dataSource.DeleteNote)
		router.Post("/notes", dataSource.CreateNote)
		router.Patch("/notes", dataSource.UpdateNote)
		router.Get("/notes", dataSource.GetAllNotes)
		// HEALTH
		router.Get("/health", handler.CheckHealth)
	})

	// START SERVER
	log.Fatal(app.Listen("127.0.0.1:8080"))
}
