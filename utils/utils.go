package utils

import (
	"e-commerce-api/app/models"
	"errors"
	"regexp"
	"strings"

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

func CleanString(s string) string {
	removedChar := regexp.MustCompile(`\n+\s+`)
	s = removedChar.ReplaceAllString(s, " ")
	return strings.Trim(s, " ")
}

func Response(c *fiber.Ctx, resp interface{}) error {
	switch resp.(type) {
	case models.ErrorMessage:
		return models.ErrorResponse(c, resp.(models.ErrorMessage))
	case models.SuccessMessage:
		return models.SuccessResponse(c, resp.(models.SuccessMessage))
	}
	return nil
}
