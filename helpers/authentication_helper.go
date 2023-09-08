package helpers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RecoverToken(c *fiber.Ctx) int32 {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	idInt := int32(id)
	return idInt
}

func CreateToken(id int32) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 60000).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
