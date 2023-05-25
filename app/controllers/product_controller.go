package controllers

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
	"errors"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return ProductController{ProductService: productService}
}

func (this *ProductController) Get(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.ErrorResponse(c, models.ErrorMessage{Err: errors.New("product not found"), StatusCode: 404})
	}

	product, custErr := this.ProductService.Get(productID)
	if custErr.Err != nil {
		return models.ErrorResponse(c, custErr)
	}

	return models.SuccessResponse(c, models.SuccessMessage{Message: product, StatusCode: 200})
}

func (this *ProductController) GetAll(c *fiber.Ctx) error {
	product, err := this.ProductService.GetAll()
	if err.Err != nil {
		return models.ErrorResponse(c, err)
	}

	return models.SuccessResponse(c, models.SuccessMessage{Message: product, StatusCode: 200})
}

func (this *ProductController) GetMultiple(c *fiber.Ctx) error {
	type IDS struct {
		IDs []int `json:"ids"`
	}

	var ids IDS

	if err := c.BodyParser(&ids); err != nil {
		return models.ErrorResponse(c, models.ErrorMessage{Err: err, StatusCode: 400})
	}

	product, err := this.ProductService.GetMultiple(ids.IDs)
	if err.Err != nil {
		return models.ErrorResponse(c, err)
	}

	return models.SuccessResponse(c, models.SuccessMessage{Message: product, StatusCode: 200})
}
