package repositories

import (
	"e-commerce-api/app/models"

	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	Get(productID int) (models.Product, error, int)
	GetAll() ([]models.Product, error, int)
	GetMultiple([]int) ([]models.Product, error, int)
}

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &productRepository{db: db}
}

func (this *productRepository) Get(productID int) (models.Product, error, int) {
	var product models.Product
	sqlStmt := fmt.Sprintf(`SELECT id, name, price, description, image FROM product WHERE id = '%d'`, productID)

	err := this.db.QueryRow(context.Background(),
		sqlStmt).Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return models.Product{}, errors.New("product not found"), 404
		}
		return models.Product{}, err, 500
	}

	if product.ID == 0 {
		return models.Product{}, errors.New("product doesn't exists"), 404
	}

	return product, nil, 0
}

func (this *productRepository) GetAll() ([]models.Product, error, int) {
	sqlStmt := fmt.Sprintf(`SELECT id, name, price, description, image FROM product`)

	rows, err := this.db.Query(context.Background(), sqlStmt)
	defer rows.Close()
	if err != nil {
		return []models.Product{}, err, 500
	}

	var results []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
		if err != nil {
			return []models.Product{}, err, 500
		}
		results = append(results, product)
	}
	return results, nil, 0
}

func (this *productRepository) GetMultiple(ids []int) ([]models.Product, error, int) {
	sqlStmt := `SELECT id, name, price, description, image FROM product WHERE id = ANY($1::integer[])`

	rows, err := this.db.Query(context.Background(), sqlStmt, ids)
	if err != nil {
		return []models.Product{}, err, 500
	}

	var results []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Image)
		if err != nil {
			return []models.Product{}, err, 500
		}
		results = append(results, product)
	}
	return results, nil, 0
}
