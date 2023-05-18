package repositories

import (
	"e-commerce-api/app/models"

	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository interface {
	SignUp(customer models.Customer) models.Result
	IsExists(field, value string) bool
	SignIn(models.SignInRequest) models.Result
}

type customerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) CustomerRepository {
	return &customerRepository{db: db}
}

func (this *customerRepository) SignUp(customer models.Customer) models.Result {
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
		return models.Result{
			Data:       models.Customer{},
			Err:        err,
			StatusCode: 500}
	}

	customer.ID = customerID
	return models.Result{
		Data:       customer,
		Err:        nil,
		StatusCode: 0}
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

func (this *customerRepository) SignIn(customer models.SignInRequest) models.Result {
	var cust models.Customer

	sqlStmt := fmt.Sprintf("SELECT id, name, email, password FROM customer WHERE email = '%s'",
		customer.Email)

	err := this.db.QueryRow(context.Background(), sqlStmt).Scan(
		&cust.ID, &cust.Name, &cust.Email, &cust.Password)

	if err != nil {
		return models.Result{
			Data:       models.Customer{},
			Err:        err,
			StatusCode: 404,
		}
	}

	return models.Result{
		Data:       cust,
		Err:        nil,
		StatusCode: 0,
	}
}
