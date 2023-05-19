package controllers

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	CustomerService services.CustomerService
}

func (this *CustomerController) SingUp(c *fiber.Ctx) error {
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return err
	}

	resp, err := this.CustomerService.SignUp(customer)
	if err.Err != nil {
		errResp := models.ErrorResponse{
			Message:    err.Err.Error(),
			StatusCode: err.StatusCode}
		return errResp.Resp(c)
	}

	successResp := models.SuccessResponse{
		Message:    models.SignUpResponseOk{Token: resp},
		StatusCode: 201,
	}
	return successResp.Resp(c)
}

func (this *CustomerController) SignIn(c *fiber.Ctx) error {
	var customer models.SignInRequest
	if err := c.BodyParser(&customer); err != nil {
		return err
	}

	resp, err := this.CustomerService.SignIn(customer)
	if err.Err != nil {
		errResp := models.ErrorResponse{
			Message:    err.Err.Error(),
			StatusCode: err.StatusCode,
		}
		return errResp.Resp(c)
	}

	successResp := models.SuccessResponse{
		Message:    models.SignUpResponseOk{Token: resp},
		StatusCode: 201,
	}
	return successResp.Resp(c)
}
