package restapi

import (
	"e-commerce-api/app/controllers"
	"e-commerce-api/app/models"
	"e-commerce-api/utils"

	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	customerController controllers.CustomerController
}

func NewCustomerHandler(controller controllers.CustomerController) customerHandler {
	return customerHandler{customerController: controller}
}

func (this *customerHandler) SignUp(c *fiber.Ctx) error {
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		resp := models.ErrorResponse(c, models.ErrorMessage{Err: err, StatusCode: 400})
		return utils.Response(c, resp)
	}

	resp := this.customerController.SingUp(customer)
	return utils.Response(c, resp)
}

func (this *customerHandler) SignIn(c *fiber.Ctx) error {
	var customer models.SignInRequest
	if err := c.BodyParser(&customer); err != nil {
		resp := models.ErrorResponse(c, models.ErrorMessage{Err: err, StatusCode: 400})
		return utils.Response(c, resp)
	}

	resp := this.customerController.SignIn(customer)
	return utils.Response(c, resp)
}
