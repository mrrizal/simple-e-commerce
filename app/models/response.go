package models

import "github.com/gofiber/fiber/v2"

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

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int
}

func (this *ErrorResponse) Resp(c *fiber.Ctx) error {
	c.Status(this.StatusCode)
	return c.JSON(this.Message)
}

type SuccessResponse struct {
	Message    interface{}
	StatusCode int
}

func (this *SuccessResponse) Resp(c *fiber.Ctx) error {
	c.Status(this.StatusCode)
	return c.JSON(this.Message)
}
