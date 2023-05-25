package repositories

import (
	"e-commerce-api/app/database"
	"e-commerce-api/app/models"

	"context"
	"fmt"
)

type CustomerRepository interface {
	SignUp(customer models.Customer) (models.Customer, models.ErrorMessage)
	IsExists(field, value string) bool
	SignIn(models.SignInRequest) (models.Customer, models.ErrorMessage)
}

type customerRepository struct {
	db database.DB
}

func NewCustomerRepository(db database.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (this *customerRepository) SignUp(customer models.Customer) (models.Customer, models.ErrorMessage) {
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
		return models.Customer{}, models.ErrorMessage{Err: err, StatusCode: 500}
	}

	customer.ID = customerID
	return customer, models.ErrorMessage{Err: nil, StatusCode: 0}
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

func (this *customerRepository) SignIn(customer models.SignInRequest) (models.Customer, models.ErrorMessage) {
	var cust models.Customer

	sqlStmt := fmt.Sprintf("SELECT id, name, email, password FROM customer WHERE email = '%s'",
		customer.Email)

	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(
		&cust.ID, &cust.Name, &cust.Email, &cust.Password)

	if err != nil {
		return models.Customer{}, models.ErrorMessage{Err: err, StatusCode: 404}
	}

	return cust, models.ErrorMessage{Err: nil, StatusCode: 0}
}
