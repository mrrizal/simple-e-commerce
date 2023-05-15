package services

import (
	"tabungan-api/app/models"
	"tabungan-api/app/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductRepository(db *pgxpool.Pool) ProductService {
	return ProductService{
		ProductRepository: repositories.NewProductRepository(db),
	}
}

func (this *ProductService) Get(productID int) (models.Product, error, int) {
	product, err, statusCode := this.ProductRepository.Get(productID)
	if err != nil {
		return models.Product{}, err, statusCode
	}
	return product, nil, 0
}
