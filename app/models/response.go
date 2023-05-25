package models

import (
	"github.com/gofiber/fiber/v2"
)

type SignUpResponseOk struct {
	Token string `json:"token"`
}

type ProductResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type OrderResp struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type OrderDataResp struct {
	Results []OrderData `json:"results"`
}

type Result struct {
	Data       interface{}
	Err        error
	StatusCode int
}

type ErrResp struct {
	Message string `json:"message"`
}

func ErrorResponse(c *fiber.Ctx, err ErrorMessage) error {
	errResp := ErrResp{Message: err.Err.Error()}
	c.Status(err.StatusCode)
	return c.JSON(errResp)
}

type SuccessMessage struct {
	Message    interface{}
	StatusCode int
}

func SuccessResponse(c *fiber.Ctx, message SuccessMessage) error {
	c.Status(message.StatusCode)
	return c.JSON(message.Message)
}

type ErrorMessage struct {
	Err        error
	StatusCode int
}

type CustomerRepositoryResult struct {
	Data       Customer
	Err        error
	StatusCode int
}

type TokenResponse struct {
	Data       string
	Err        error
	StatusCode int
}

type OrderRepositoryOrderResult struct {
	Data       OrderResp
	Err        error
	StatusCode int
}
