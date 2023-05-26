package services

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/repositories"
)

type ProductService interface {
	Get(int) (models.ProductResp, models.ErrorMessage)
	GetAll() (models.ProductList, models.ErrorMessage)
	GetMultiple(ids []int) (models.ProductList, models.ErrorMessage)
}

type productService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) ProductService {
	return &productService{ProductRepository: repository}
}

func (this *productService) Parse(product models.Product) models.ProductResp {
	var result models.ProductResp
	result.ID = product.ID
	result.Name = product.Name
	result.Price = product.Price
	result.Description = product.Description.String
	result.Image = product.Image.String
	return result
}

func (this *productService) Get(productID int) (models.ProductResp, models.ErrorMessage) {
	product, err := this.ProductRepository.Get(productID)
	if err.Err != nil {
		return models.ProductResp{}, err
	}

	result := this.Parse(product)
	return result, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *productService) GetAll() (models.ProductList, models.ErrorMessage) {
	products, err := this.ProductRepository.GetAll()
	if err.Err != nil {
		return models.ProductList{}, err
	}

	var result models.ProductList
	for _, product := range products {
		result.Results = append(result.Results, this.Parse(product))
	}
	return result, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *productService) GetMultiple(ids []int) (models.ProductList, models.ErrorMessage) {
	products, err := this.ProductRepository.GetMultiple(ids)
	if err.Err != nil {
		return models.ProductList{}, err
	}

	var result models.ProductList
	for _, product := range products {
		result.Results = append(result.Results, this.Parse(product))
	}
	return result, models.ErrorMessage{Err: nil, StatusCode: 0}
}
