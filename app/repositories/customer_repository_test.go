package repositories

import (
	"context"
	"e-commerce-api/app/database"
	"e-commerce-api/app/models"
	"e-commerce-api/utils"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	// mock database
	db := database.NewMockDB()

	customerPayload := models.Customer{Name: "rizal", Email: "rizal@test.com", Password: "qwerty"}

	t.Run("sign-up success", func(t *testing.T) {
		// customize QueryRow function, we want test this to make we do query correctly
		db.MockQueryRow = func(ctx context.Context, sql string, args ...any) database.Row {
			sql = utils.CleanString(sql)
			assert.Equal(t, "INSERT INTO Customer (name, email, password) VALUES ($1, $2, $3) RETURNING id", sql)
			assert.Equal(t, customerPayload.Name, args[0])
			assert.Equal(t, customerPayload.Email, args[1])
			assert.Equal(t, customerPayload.Password, args[2])
			return &db.MockRow
		}

		// do sign-up test
		repository := NewCustomerRepository(&db)
		_, err := repository.SignUp(customerPayload)
		assert.Equal(t, models.ErrorMessage{Err: nil, StatusCode: 0}, err)
	})

	t.Run("sign-up error", func(t *testing.T) {
		// mock row
		row := database.NewMockRow()

		// update Scan function of Row, we want this function return error
		row.MockScan = func(dest ...any) error {
			return errors.New("error")
		}
		// inject the new row with updated scan function to the db (mock db)
		db.MockRow = row

		// do sign-up test, we expect it's will raise error
		repository := NewCustomerRepository(&db)
		_, err := repository.SignUp(customerPayload)
		assert.Equal(t, models.ErrorMessage{Err: errors.New("error"), StatusCode: 500}, err)
	})

}
