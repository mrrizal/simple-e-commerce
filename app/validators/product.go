package validators

import (
	"e-commerce-api/app/services"
	"errors"
)

type ProductValidator struct {
	ProductService services.ProductService
}

func (this *ProductValidator) ValidateProducts(ids []int) ([]int, error, int) {
	products, err, statusCode := this.ProductService.GetMultiple(ids)
	if err != nil {
		return []int{}, err, statusCode
	}

	result := []int{}
	for _, product := range products.Results {
		result = append(result, product.ID)
	}

	if len(result) == 0 {
		return result, errors.New("invalid products id"), 400
	}

	return result, nil, 0
}
