package repositories

import (
	"context"
	"errors"
	"fmt"
	"tabungan-api/app/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	Get(productID int) (models.Product, error, int)
	// GetAll() (models.ProductList, error, int)
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

	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(&product)
	if err != nil {
		return models.Product{}, err, 500
	}

	if product.ID == 0 {
		return models.Product{}, errors.New("product doesn't exists"), 404
	}

	return product, nil, 0
}

// func (this *customerRepository) IsExists(field, value string) bool {
// 	var customerID int

// 	sqlStmt := fmt.Sprintf("SELECT id FROM customer WHERE %s = '%s'", field, value)
// 	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(&customerID)
// 	if err != nil {
// 		return false
// 	}
// 	return customerID != 0
// }

// func (this *customerRepository) SignIn(customer models.SignInRequest) models.Customer {
// 	var cust models.Customer

// 	sqlStmt := fmt.Sprintf("SELECT id, name, email, password FROM customer WHERE email = '%s'",
// 		customer.Email)

// 	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(
// 		&cust.ID, &cust.Name, &cust.Email, &cust.Password)

// 	if err != nil {
// 		return models.Customer{}
// 	}

// 	return cust
// }
