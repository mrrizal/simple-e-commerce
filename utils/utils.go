package utils

import (
	"e-commerce-api/app/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ParseJWTToken(tokenString, secretKey string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return map[string]interface{}{}, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return fiber.Map{"claims": claims}, nil
	}

	return map[string]interface{}{}, errors.New("invalid token")

}

func ErrorResp(c *fiber.Ctx, message string, statusCode int) error {
	var result models.ErrorResponse
	result.Message = message
	if statusCode == 0 {
		statusCode = 400
	}
	c.Status(statusCode)
	return c.JSON(result)
}
