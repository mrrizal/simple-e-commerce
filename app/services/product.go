package services

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService interface {
	Get(int) (models.ProductResp, models.CustomError)
	GetAll() (models.ProductList, models.CustomError)
	GetMultiple(ids []int) (models.ProductList, models.CustomError)
}

type productService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(db *pgxpool.Pool) ProductService {
	return &productService{
		ProductRepository: repositories.NewProductRepository(db),
	}
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

func (this *productService) Get(productID int) (models.ProductResp, models.CustomError) {
	product, err := this.ProductRepository.Get(productID)
	if err.Err != nil {
		return models.ProductResp{}, err
	}

	result := this.Parse(product)
	return result, models.CustomError{Err: nil, StatusCode: 0}
}

func (this *productService) GetAll() (models.ProductList, models.CustomError) {
	products, err := this.ProductRepository.GetAll()
	if err.Err != nil {
		return models.ProductList{}, err
	}

	var result models.ProductList
	for _, product := range products {
		result.Results = append(result.Results, this.Parse(product))
	}
	return result, models.CustomError{Err: nil, StatusCode: 0}
}

func (this *productService) GetMultiple(ids []int) (models.ProductList, models.CustomError) {
	products, err := this.ProductRepository.GetMultiple(ids)
	if err.Err != nil {
		return models.ProductList{}, err
	}

	var result models.ProductList
	for _, product := range products {
		result.Results = append(result.Results, this.Parse(product))
	}
	return result, models.CustomError{Err: nil, StatusCode: 0}
}
