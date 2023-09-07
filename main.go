package main

import (
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/multi-user-notes-app/db/internals"
	handler "github.com/multi-user-notes-app/handlers"
	"github.com/multi-user-notes-app/helpers"
)

func main() {
	app := fiber.New()

	helpers.InitEnv()
	helpers.InitValidator()
	internals.InitDB()

	app.Use(logger.New())
	app.Use("/api", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
	}))

	app.Route("/authenticate", func(router fiber.Router) {
		router.Post("/login", handler.Login)
		router.Post("/register", handler.Register)
	})

	app.Route("/api", func(router fiber.Router) {
		router.Route("/user", func(router fiber.Router) {
			router.Delete("/", handler.DeleteUserById)
		})
		router.Route("/notes", func(router fiber.Router) {
			router.Get("/favourites", handler.ListFavourites)
			router.Get("/:id", handler.GetNoteById)
			router.Delete("/:id", handler.DeleteNote)
			router.Post("/", handler.CreateNote)
			router.Patch("/", handler.UpdateNote)
			router.Get("/", handler.GetAllNotes)
			router.Put("/add-to-favourites/:id", handler.AddToFavourites)
			router.Put("/remove-to-favourites/:id", handler.RemoveFromFavourite)
		})

		router.Get("/health", handler.CheckHealth)
	})

	log.Fatal(app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}
