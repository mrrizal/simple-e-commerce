package validators

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
	"errors"
)

type ProductValidator interface {
	ValidateProducts([]int) ([]int, models.ErrorMessage)
}

type productValidator struct {
	ProductService services.ProductService
}

func NewProductValidator(service services.ProductService) ProductValidator {
	return &productValidator{ProductService: service}
}

func (this *productValidator) ValidateProducts(ids []int) ([]int, models.ErrorMessage) {
	products, err := this.ProductService.GetMultiple(ids)
	if err.Err != nil {
		return []int{}, err
	}

	result := []int{}
	for _, product := range products.Results {
		result = append(result, product.ID)
	}

	if len(result) == 0 {
		return result, models.ErrorMessage{Err: errors.New("invalid products id"), StatusCode: 400}
	}

	return result, models.ErrorMessage{Err: nil, StatusCode: 0}
}
