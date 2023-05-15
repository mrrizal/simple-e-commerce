package repositories

import (
	"context"
	"fmt"
	"tabungan-api/app/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository interface {
	SignUp(customer models.Customer) (models.Customer, error, int)
	IsExists(field, value string) bool
	SignIn(models.SignInRequest) models.Customer
}

type customerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) CustomerRepository {
	return &customerRepository{db: db}
}

func (this *customerRepository) SignUp(customer models.Customer) (models.Customer, error, int) {
	sqlStmt := `
		INSERT INTO Customer (name, email, password)
		VALUES ($1, $2, $3) RETURNING id
	`
	var customerID int
	err := this.db.QueryRow(
		context.Background(),
		sqlStmt,
		customer.Name,
		customer.Email,
		customer.Password).Scan(&customerID)

	if err != nil {
		return models.Customer{}, err, 500
	}

	customer.ID = customerID
	return customer, nil, 0
}

func (this *customerRepository) IsExists(field, value string) bool {
	var customerID int

	sqlStmt := fmt.Sprintf("SELECT id FROM customer WHERE %s = '%s'", field, value)
	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(&customerID)
	if err != nil {
		return false
	}
	return customerID != 0
}

func (this *customerRepository) SignIn(customer models.SignInRequest) models.Customer {
	var cust models.Customer

	sqlStmt := fmt.Sprintf("SELECT id, name, email, password FROM customer WHERE email = '%s'",
		customer.Email)

	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(
		&cust.ID, &cust.Name, &cust.Email, &cust.Password)

	if err != nil {
		return models.Customer{}
	}

	return cust
}
