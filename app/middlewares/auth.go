package middlewares

import (
	"e-commerce-api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	temp := c.Get("Authorization")
	if !strings.HasPrefix(temp, "Bearer") {
		return utils.ErrorResp(c, "invalid token", 401)
	}

	if len(strings.Split(temp, " ")) != 2 {
		return utils.ErrorResp(c, "invalid token", 401)
	}

	return c.Next()
}
