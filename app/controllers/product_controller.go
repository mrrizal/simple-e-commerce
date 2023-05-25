package controllers

import (
	"e-commerce-api/app/services"
	"e-commerce-api/utils"

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
		utils.ErrorResp(c, "product not found", 404)
	}

	product, custErr := this.ProductService.Get(productID)
	if custErr.Err != nil {
		return utils.ErrorResp(c, custErr.Err.Error(), custErr.StatusCode)
	}

	c.Status(200)
	return c.JSON(product)
}

func (this *ProductController) GetAll(c *fiber.Ctx) error {
	product, err := this.ProductService.GetAll()
	if err.Err != nil {
		return utils.ErrorResp(c, err.Err.Error(), err.StatusCode)
	}

	c.Status(200)
	return c.JSON(product)
}

func (this *ProductController) GetMultiple(c *fiber.Ctx) error {
	type IDS struct {
		IDs []int `json:"ids"`
	}

	var ids IDS

	if err := c.BodyParser(&ids); err != nil {
		return utils.ErrorResp(c, err.Error(), 400)
	}

	product, err := this.ProductService.GetMultiple(ids.IDs)
	if err.Err != nil {
		return utils.ErrorResp(c, err.Err.Error(), err.StatusCode)
	}

	c.Status(200)
	return c.JSON(product)
}
