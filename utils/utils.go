package utils

import (
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
