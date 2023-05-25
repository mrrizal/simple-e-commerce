package controllers

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return CustomerController{customerService: customerService}
}

func (this *CustomerController) SingUp(c *fiber.Ctx) error {
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return err
	}

	resp, err := this.customerService.SignUp(customer)
	if err.Err != nil {
		return models.ErrorResponse(c, err)
	}

	return models.SuccessResponse(c, models.SuccessMessage{
		Message:    models.SignUpResponseOk{Token: resp},
		StatusCode: 201,
	})
}

func (this *CustomerController) SignIn(c *fiber.Ctx) error {
	var customer models.SignInRequest
	if err := c.BodyParser(&customer); err != nil {
		return err
	}

	resp, err := this.customerService.SignIn(customer)
	if err.Err != nil {
		return models.ErrorResponse(c, err)
	}

	return models.SuccessResponse(c, models.SuccessMessage{
		Message:    models.SignUpResponseOk{Token: resp},
		StatusCode: 200,
	})
}
