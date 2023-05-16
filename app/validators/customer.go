package validators

import (
	"e-commerce-api/app/configs"
	"e-commerce-api/app/services"
	"e-commerce-api/utils"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type CustomerValidator struct {
	CustomerService services.CustomerService
	Config          configs.Config
}

func (this *CustomerValidator) ValidateCustomer(token string) (int, error, int) {
	customerInfo, err := utils.ParseJWTToken(token, this.Config.SecretKey)
	if err != nil {
		return 0, err, 400
	}

	customerID := customerInfo["claims"].(jwt.MapClaims)["id"].(float64)
	if !this.CustomerService.IsExists(int(customerID)) {
		return 0, errors.New("invalid token"), 401
	}
	return int(customerID), nil, 0
}
