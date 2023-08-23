package handler

import (
	"github.com/Fiber-CRUD/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// func Login(c *fiber.Ctx) error {

// 	loginForm := new(forms.Login)

// 	if err := c.BodyParser(loginForm); err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	user, err := db.GetUser(*loginForm)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	if isValid := checkPasswordHash(loginForm.Password, user.Password); !isValid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
// 			"message": "Invalid password",
// 		})
// 	}

// 	claims := jwt.MapClaims{
// 		"id":    user.Id,
// 		"email": user.Email,
// 		"exp":   time.Now().Add(time.Minute * 60).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	return c.JSON(fiber.Map{"token": t})
// }

func Register(c *fiber.Ctx) error {

	createUserParams := new(db.CreateUserParams)

	if err := c.BodyParser(createUserParams); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	password, err := hashPassword(createUserParams.Password)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	createUserParams.Password = password

	user, err := db.DBConnection.DB.CreateUser(c.Context(), *createUserParams)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"user": user,
	})
}

func recoverToken(c *fiber.Ctx) (float64, string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	email := claims["email"].(string)
	return id, email
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
