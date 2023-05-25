package repositories

import (
	"e-commerce-api/app/database"
	"e-commerce-api/app/models"

	"context"
	"errors"
	"fmt"
)

type ProductRepository interface {
	Get(productID int) (models.Product, models.ErrorMessage)
	GetAll() ([]models.Product, models.ErrorMessage)
	GetMultiple([]int) ([]models.Product, models.ErrorMessage)
}

type productRepository struct {
	db database.DB
}

func NewProductRepository(db database.DB) ProductRepository {
	return &productRepository{db: db}
}

func (this *productRepository) Get(productID int) (models.Product, models.ErrorMessage) {
	var product models.Product
	sqlStmt := fmt.Sprintf(`SELECT id, name, price, description, image FROM product WHERE id = '%d'`, productID)

	err := this.db.QueryRow(context.Background(),
		sqlStmt).Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return models.Product{}, models.ErrorMessage{Err: errors.New("product not found"), StatusCode: 404}
		}
		return models.Product{}, models.ErrorMessage{Err: err, StatusCode: 500}
	}

	if product.ID == 0 {
		return models.Product{}, models.ErrorMessage{Err: errors.New("product doesn't exists"), StatusCode: 404}
	}

	return product, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *productRepository) GetAll() ([]models.Product, models.ErrorMessage) {
	sqlStmt := fmt.Sprintf(`SELECT id, name, price, description, image FROM product`)

	rows, err := this.db.Query(context.Background(), sqlStmt)
	defer rows.Close()
	if err != nil {
		return []models.Product{}, models.ErrorMessage{Err: err, StatusCode: 500}
	}

	var results []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
		if err != nil {
			return []models.Product{}, models.ErrorMessage{Err: err, StatusCode: 500}
		}
		results = append(results, product)
	}
	return results, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *productRepository) GetMultiple(ids []int) ([]models.Product, models.ErrorMessage) {
	sqlStmt := `SELECT id, name, price, description, image FROM product WHERE id = ANY($1::integer[])`

	rows, err := this.db.Query(context.Background(), sqlStmt, ids)
	if err != nil {
		return []models.Product{}, models.ErrorMessage{Err: err, StatusCode: 500}
	}

	var results []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
		if err != nil {
			return []models.Product{}, models.ErrorMessage{Err: err, StatusCode: 500}
		}
		results = append(results, product)
	}
	return results, models.ErrorMessage{Err: nil, StatusCode: 0}
}
