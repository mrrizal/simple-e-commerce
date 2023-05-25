package repositories

import (
	"e-commerce-api/app/models"

	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	Get(productID int) (models.Product, models.CustomError)
	GetAll() ([]models.Product, models.CustomError)
	GetMultiple([]int) ([]models.Product, models.CustomError)
}

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &productRepository{db: db}
}

func (this *productRepository) Get(productID int) (models.Product, models.CustomError) {
	var product models.Product
	sqlStmt := fmt.Sprintf(`SELECT id, name, price, description, image FROM product WHERE id = '%d'`, productID)

	err := this.db.QueryRow(context.Background(),
		sqlStmt).Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return models.Product{}, models.CustomError{Err: errors.New("product not found"), StatusCode: 404}
		}
		return models.Product{}, models.CustomError{Err: err, StatusCode: 500}
	}

	if product.ID == 0 {
		return models.Product{}, models.CustomError{Err: errors.New("product doesn't exists"), StatusCode: 404}
	}

	return product, models.CustomError{Err: nil, StatusCode: 0}
}

func (this *productRepository) GetAll() ([]models.Product, models.CustomError) {
	sqlStmt := fmt.Sprintf(`SELECT id, name, price, description, image FROM product`)

	rows, err := this.db.Query(context.Background(), sqlStmt)
	defer rows.Close()
	if err != nil {
		return []models.Product{}, models.CustomError{Err: err, StatusCode: 500}
	}

	var results []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
		if err != nil {
			return []models.Product{}, models.CustomError{Err: err, StatusCode: 500}
		}
		results = append(results, product)
	}
	return results, models.CustomError{Err: nil, StatusCode: 0}
}

func (this *productRepository) GetMultiple(ids []int) ([]models.Product, models.CustomError) {
	sqlStmt := `SELECT id, name, price, description, image FROM product WHERE id = ANY($1::integer[])`

	rows, err := this.db.Query(context.Background(), sqlStmt, ids)
	if err != nil {
		return []models.Product{}, models.CustomError{Err: err, StatusCode: 500}
	}

	var results []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
		if err != nil {
			return []models.Product{}, models.CustomError{Err: err, StatusCode: 500}
		}
		results = append(results, product)
	}
	return results, models.CustomError{Err: nil, StatusCode: 0}
}
