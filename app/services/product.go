package services

import (
	"tabungan-api/app/models"
	"tabungan-api/app/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(db *pgxpool.Pool) ProductService {
	return ProductService{
		ProductRepository: repositories.NewProductRepository(db),
	}
}

func (this *ProductService) Parse(product models.Product) models.ProductResp {
	var result models.ProductResp
	result.ID = product.ID
	result.Name = product.Name
	result.Price = product.Price
	result.Description = product.Description.String
	result.Image = product.Image.String
	return result
}

func (this *ProductService) Get(productID int) (models.ProductResp, error, int) {
	product, err, statusCode := this.ProductRepository.Get(productID)
	if err != nil {
		return models.ProductResp{}, err, statusCode
	}

	result := this.Parse(product)
	return result, nil, 0
}

func (this *ProductService) GetAll() (models.ProductList, error, int) {
	products, err, statusCode := this.ProductRepository.GetAll()
	if err != nil {
		return models.ProductList{}, err, statusCode
	}

	var result models.ProductList
	for _, product := range products {
		result.Results = append(result.Results, this.Parse(product))
	}
	return result, nil, 0
}

func (this *ProductService) GetMultiple(ids []int) (models.ProductList, error, int) {
	products, err, statusCode := this.ProductRepository.GetMultiple(ids)
	if err != nil {
		return models.ProductList{}, err, statusCode
	}

	var result models.ProductList
	for _, product := range products {
		result.Results = append(result.Results, this.Parse(product))
	}
	return result, nil, 0
}
