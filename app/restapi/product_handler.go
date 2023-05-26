// fiber handler
package restapi

import (
	"e-commerce-api/app/controllers"
	"e-commerce-api/app/models"
	"e-commerce-api/utils"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productController controllers.ProductController
}

func NewProductHandler(controller controllers.ProductController) productHandler {
	return productHandler{productController: controller}
}

func (this *productHandler) Get(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Response(c, models.ErrorMessage{Err: errors.New("product not found"), StatusCode: 404})
	}

	resp := this.productController.Get(productID)
	return utils.Response(c, resp)
}

func (this *productHandler) GetAll(c *fiber.Ctx) error {
	resp := this.productController.GetAll()
	return utils.Response(c, resp)
}

func (this *productHandler) GetMultiple(c *fiber.Ctx) error {
	type IDS struct {
		IDs []int `json:"ids"`
	}

	var ids IDS

	if err := c.BodyParser(&ids); err != nil {
		return utils.Response(c, models.ErrorMessage{Err: err, StatusCode: 400})
	}

	resp := this.productController.GetMultiple(ids.IDs)
	return utils.Response(c, resp)
}
