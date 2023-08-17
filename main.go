package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	handler "github.com/Fiber-CRUD/controllers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
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
	dataSource, err := openDB()
	if err != nil {
		log.Fatal("ERROR CONNECTING TO DB")
	}

	app.Route("/authenticate", func(router fiber.Router) {
		router.Post("/login", login)
		router.Get("/", accessible)
		router.Get("/restricted", restricted)
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

func login(c *fiber.Ctx) error {
	println("LOGIN")
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	fmt.Println("1")
	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	fmt.Println("2")
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	fmt.Println("3")
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) error {
	fmt.Println("ACC")
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	fmt.Println("RES")
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
