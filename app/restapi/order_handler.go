// fiber handler
package restapi

import (
	"e-commerce-api/app/controllers"
	"e-commerce-api/app/models"
	"e-commerce-api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderController controllers.OrderController
}

func NewOrderHandler(controller controllers.OrderController) orderHandler {
	return orderHandler{orderController: controller}
}

func (this *orderHandler) CreateOrder(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]

	// parse body request
	var order models.OrderRequest
	if err := c.BodyParser(&order); err != nil {
		errMessage := models.ErrorMessage{Err: err, StatusCode: 400}
		return utils.Response(c, errMessage)
	}

	resp := this.orderController.CreateOrder(token, order)
	return utils.Response(c, resp)
}

func (this *orderHandler) Get(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]
	resp := this.orderController.Get(token)
	return utils.Response(c, resp)
}
