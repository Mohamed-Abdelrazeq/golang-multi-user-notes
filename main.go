package main

import (
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/multi-user-notes-app/db"
	handler "github.com/multi-user-notes-app/handlers"
	"github.com/multi-user-notes-app/helpers"
	supa "github.com/multi-user-notes-app/supabase"
)

func main() {
	app := fiber.New()
	helpers.LoadEnv()

	app.Use(logger.New())
	app.Use("/api", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
	}))

	supa.OpenClientConnection()
	err := db.OpenDBConnection()
	if err != nil {
		log.Fatal("ERROR CONNECTING TO DB")
	}

	app.Route("/authenticate", func(router fiber.Router) {
		router.Post("/authenticate-user", handler.AuthenticateUser)
		// router.Post("/create-user", handler.CreateUser)
		router.Post("/create-user", supa.CreateUser)
	})

	app.Route("/api", func(router fiber.Router) {
		// NOTES
		router.Get("/notes/favourites", handler.ListFavourites)
		router.Get("/notes/:id", handler.GetNoteById)
		router.Delete("/notes/:id", handler.DeleteNote)
		router.Post("/notes", handler.CreateNote)
		router.Patch("/notes", handler.UpdateNote)
		router.Get("/notes", handler.GetAllNotes)
		router.Put("/notes/add-to-favourites/:id", handler.AddToFavourites)
		router.Put("/notes/remove-to-favourites/:id", handler.RemoveFromFavourite)
		// HEALTH
		router.Get("/health", handler.CheckHealth)
	})

	log.Fatal(app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}
