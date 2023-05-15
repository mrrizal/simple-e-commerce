package controllers

import (
	"tabungan-api/app/models"
	"tabungan-api/app/services"
	"tabungan-api/utils"

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

	token, err, statusCode := this.CustomerService.SignUp(customer)
	if err != nil {
		return utils.ErrorResp(c, err.Error(), statusCode)
	}

	var result models.SignUpResponseOk
	result.Token = token

	c.Status(200)
	return c.JSON(result)
}

func (this *CustomerController) SignIn(c *fiber.Ctx) error {
	var customer models.SignInRequest
	if err := c.BodyParser(&customer); err != nil {
		return err
	}

	token, err, statusCode := this.CustomerService.SignIn(customer)
	if err != nil {
		return utils.ErrorResp(c, err.Error(), statusCode)
	}

	var result models.SignUpResponseOk
	result.Token = token

	c.Status(200)
	return c.JSON(result)
}
