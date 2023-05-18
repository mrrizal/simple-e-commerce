package controllers

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
	"e-commerce-api/utils"

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

	resp := this.CustomerService.SignUp(customer)
	if resp.Err != nil {
		errResp := utils.ErrorResponse{
			Message:    resp.Err.Error(),
			StatusCode: resp.StatusCode}
		return errResp.Resp(c)
	}

	successResp := utils.SuccessResponse{
		Message:    models.SignUpResponseOk{Token: resp.Data.(string)},
		StatusCode: 201,
	}
	return successResp.Resp(c)
}

func (this *CustomerController) SignIn(c *fiber.Ctx) error {
	var customer models.SignInRequest
	if err := c.BodyParser(&customer); err != nil {
		return err
	}

	resp := this.CustomerService.SignIn(customer)
	if resp.Err != nil {
		errResp := utils.ErrorResponse{
			Message:    resp.Err.Error(),
			StatusCode: resp.StatusCode,
		}
		return errResp.Resp(c)
	}

	successResp := utils.SuccessResponse{
		Message:    models.SignUpResponseOk{Token: resp.Data.(string)},
		StatusCode: 201,
	}
	return successResp.Resp(c)
}
