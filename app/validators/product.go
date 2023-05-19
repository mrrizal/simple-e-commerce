package validators

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
	"errors"
)

type ProductValidator interface {
	ValidateProducts([]int) ([]int, models.CustomError)
}

type productValidator struct {
	ProductService services.ProductService
}

func NewProductValidator(service services.ProductService) ProductValidator {
	return &productValidator{ProductService: service}
}

func (this *productValidator) ValidateProducts(ids []int) ([]int, models.CustomError) {
	products, err, statusCode := this.ProductService.GetMultiple(ids)
	if err != nil {
		return []int{}, models.CustomError{Err: err, StatusCode: statusCode}
	}

	result := []int{}
	for _, product := range products.Results {
		result = append(result, product.ID)
	}

	if len(result) == 0 {
		return result, models.CustomError{Err: errors.New("invalid products id"), StatusCode: 400}
	}

	return result, models.CustomError{Err: nil, StatusCode: 0}
}
