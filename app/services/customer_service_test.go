package services

import (
	"e-commerce-api/app/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCustomerRepository struct {
	signUp   func(customer models.Customer) (models.Customer, models.ErrorMessage)
	isExists func(field, value string) bool
	signIn   func(models.SignInRequest) (models.Customer, models.ErrorMessage)
}

func (this *mockCustomerRepository) SignUp(customer models.Customer) (models.Customer, models.ErrorMessage) {
	return this.signUp(customer)
}

func (this *mockCustomerRepository) IsExists(field, value string) bool {
	return this.isExists(field, value)
}

func (this *mockCustomerRepository) SignIn(request models.SignInRequest) (models.Customer, models.ErrorMessage) {
	return this.signIn(request)
}

func NewMockCustomerRepository() mockCustomerRepository {
	return mockCustomerRepository{
		signUp: func(customer models.Customer) (models.Customer, models.ErrorMessage) {
			return models.Customer{}, models.ErrorMessage{}
		},
		isExists: func(field, value string) bool {
			return false
		},
		signIn: func(sir models.SignInRequest) (models.Customer, models.ErrorMessage) {
			return models.Customer{}, models.ErrorMessage{}
		},
	}
}

func TestSingUp(t *testing.T) {
	t.Run("sign-up success", func(t *testing.T) {
		mockRepository := NewMockCustomerRepository()
		svc := NewCustomerService(&mockRepository)
		token, err := svc.SignUp(models.Customer{})
		assert.NotEqual(t, "", token)
		assert.Equal(t, models.ErrorMessage{Err: nil, StatusCode: 0}, err)
	})

	t.Run("sign-up email exists", func(t *testing.T) {
		mockRepository := NewMockCustomerRepository()
		mockRepository.isExists = func(field, value string) bool {
			return true
		}

		svc := NewCustomerService(&mockRepository)
		token, err := svc.SignUp(models.Customer{})
		assert.Equal(t, "", token)
		assert.Equal(t, models.ErrorMessage{Err: errors.New("email already exists"), StatusCode: 400}, err)
	})

	t.Run("sign-up error", func(t *testing.T) {
		mockRepository := NewMockCustomerRepository()
		mockRepository.signUp = func(customer models.Customer) (models.Customer, models.ErrorMessage) {
			return models.Customer{}, models.ErrorMessage{Err: errors.New("error"), StatusCode: 500}
		}

		svc := NewCustomerService(&mockRepository)
		token, err := svc.SignUp(models.Customer{})
		assert.Equal(t, "", token)
		assert.Equal(t, models.ErrorMessage{Err: errors.New("error"), StatusCode: 500}, err)
	})

}
