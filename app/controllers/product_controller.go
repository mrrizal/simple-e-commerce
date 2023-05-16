package controllers

import (
	"strconv"
	"tabungan-api/app/services"
	"tabungan-api/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ProductService services.ProductService
}

func (this *ProductController) Get(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		utils.ErrorResp(c, "product not found", 404)
	}

	product, err, statusCode := this.ProductService.Get(productID)
	if err != nil {
		return utils.ErrorResp(c, err.Error(), statusCode)
	}

	c.Status(200)
	return c.JSON(product)
}

func (this *ProductController) GetAll(c *fiber.Ctx) error {
	product, err, statusCode := this.ProductService.GetAll()
	if err != nil {
		return utils.ErrorResp(c, err.Error(), statusCode)
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

	product, err, statusCode := this.ProductService.GetMultiple(ids.IDs)
	if err != nil {
		return utils.ErrorResp(c, err.Error(), statusCode)
	}

	c.Status(200)
	return c.JSON(product)
}
