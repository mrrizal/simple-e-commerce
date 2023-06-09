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
	ValidateCustomer(string) (int, models.ErrorMessage)
}

type customerValidator struct {
	CustomerService services.CustomerService
	Config          configs.Config
}

func NewCustomerValidator(service services.CustomerService) CustomerValidator {
	return &customerValidator{
		CustomerService: service,
		Config:          *configs.GetConfig()}
}

func (this *customerValidator) ValidateCustomer(token string) (int, models.ErrorMessage) {
	customerInfo, err := utils.ParseJWTToken(token, this.Config.SecretKey)
	if err != nil {
		return 0, models.ErrorMessage{Err: err, StatusCode: 400}
	}

	customerID := customerInfo["claims"].(jwt.MapClaims)["id"].(float64)
	if !this.CustomerService.IsExists(int(customerID)) {
		return 0, models.ErrorMessage{Err: errors.New("invalid token"), StatusCode: 401}
	}
	return int(customerID), models.ErrorMessage{Err: nil, StatusCode: 0}
}
