package middlewares

import (
	"e-commerce-api/app/models"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	temp := c.Get("Authorization")
	if !strings.HasPrefix(temp, "Bearer") {
		return models.ErrorResponse(c, models.ErrorMessage{Err: errors.New("invalid token"), StatusCode: 401})
	}

	if len(strings.Split(temp, " ")) != 2 {
		return models.ErrorResponse(c, models.ErrorMessage{Err: errors.New("invalid token"), StatusCode: 401})
	}

	return c.Next()
}
