// this packag handle request from user
package controllers

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return CustomerController{customerService: customerService}
}

func (this *CustomerController) SingUp(customer models.Customer) interface{} {
	resp, err := this.customerService.SignUp(customer)
	if err.Err != nil {
		return err
	}

	return models.SuccessMessage{
		Message:    models.SignUpResponseOk{Token: resp},
		StatusCode: 201,
	}
}

func (this *CustomerController) SignIn(customer models.SignInRequest) interface{} {
	resp, err := this.customerService.SignIn(customer)
	if err.Err != nil {
		return err
	}

	return models.SuccessMessage{
		Message:    models.SignUpResponseOk{Token: resp},
		StatusCode: 200,
	}
}
