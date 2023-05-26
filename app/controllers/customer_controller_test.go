package controllers

import (
	"e-commerce-api/app/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCustomerService struct {
	signUp   func(models.Customer) (string, models.ErrorMessage)
	signIn   func(models.SignInRequest) (string, models.ErrorMessage)
	isExists func(int) bool
}

func (this *mockCustomerService) SignUp(cust models.Customer) (string, models.ErrorMessage) {
	return this.signUp(cust)
}

func (this *mockCustomerService) SignIn(req models.SignInRequest) (string, models.ErrorMessage) {
	return this.signIn(req)
}

func (this *mockCustomerService) IsExists(id int) bool {
	return this.isExists(id)
}

func NewMockCustomerService() mockCustomerService {
	return mockCustomerService{
		signUp: func(c models.Customer) (string, models.ErrorMessage) {
			return "", models.ErrorMessage{}
		},
		signIn: func(sir models.SignInRequest) (string, models.ErrorMessage) {
			return "", models.ErrorMessage{}
		},
		isExists: func(i int) bool {
			return false
		},
	}
}

func TestSignIn(t *testing.T) {
	t.Run("sign-in success", func(t *testing.T) {
		svc := NewMockCustomerService()
		controller := NewCustomerController(&svc)
		resp := controller.SignIn(models.SignInRequest{})
		assert.Equal(t, models.SuccessMessage{
			Message: models.SignUpResponseOk{Token: ""}, StatusCode: 200}, resp.(models.SuccessMessage))
	})

	t.Run("sign-in failed", func(t *testing.T) {
		svc := NewMockCustomerService()
		svc.signIn = func(sir models.SignInRequest) (string, models.ErrorMessage) {
			return "", models.ErrorMessage{Err: errors.New("error"), StatusCode: 404}
		}
		controller := NewCustomerController(&svc)
		resp := controller.SignIn(models.SignInRequest{})
		assert.Equal(t, models.ErrorMessage{
			Err: errors.New("error"), StatusCode: 404}, resp.(models.ErrorMessage))
	})
}
