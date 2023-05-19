package validators

import (
	"e-commerce-api/app/configs"
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
	"e-commerce-api/utils"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type CustomerValidator interface {
	ValidateCustomer(string) (int, models.CustomError)
}

type customerValidator struct {
	CustomerService services.CustomerService
	Config          configs.Config
}

func NewCustomerValidator(service services.CustomerService, config configs.Config) CustomerValidator {
	return &customerValidator{CustomerService: service, Config: config}
}

func (this *customerValidator) ValidateCustomer(token string) (int, models.CustomError) {
	customerInfo, err := utils.ParseJWTToken(token, this.Config.SecretKey)
	if err != nil {
		return 0, models.CustomError{Err: err, StatusCode: 400}
	}

	customerID := customerInfo["claims"].(jwt.MapClaims)["id"].(float64)
	if !this.CustomerService.IsExists(int(customerID)) {
		return 0, models.CustomError{Err: errors.New("invalid token"), StatusCode: 401}
	}
	return int(customerID), models.CustomError{Err: nil, StatusCode: 0}
}
