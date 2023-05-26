package services

import (
	"e-commerce-api/app/configs"
	"e-commerce-api/app/models"
	"e-commerce-api/app/repositories"
	"strconv"

	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService interface {
	SignUp(models.Customer) (string, models.ErrorMessage)
	SignIn(models.SignInRequest) (string, models.ErrorMessage)
	IsExists(int) bool
}

type customerService struct {
	CustomerRepository repositories.CustomerRepository
	config             configs.Config
}

func NewCustomerService(repository repositories.CustomerRepository) CustomerService {
	return &customerService{
		CustomerRepository: repository,
		config:             *configs.GetConfig(),
	}
}

func (this *customerService) generatePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (this *customerService) verifyPassword(customerPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(customerPassword), []byte(password))
	return err == nil
}

func (this *customerService) generateJWTToken(customer models.Customer) (string, error) {
	claims := jwt.MapClaims{
		"id":    customer.ID,
		"name":  customer.Name,
		"email": customer.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(this.config.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (this *customerService) SignUp(customer models.Customer) (string, models.ErrorMessage) {
	emailIsExists := this.CustomerRepository.IsExists("email", customer.Email)
	if emailIsExists {
		return "", models.ErrorMessage{Err: errors.New("email already exists"), StatusCode: 400}
	}

	tempPassword, err := this.generatePassword(customer.Password)
	if err != nil {
		return "", models.ErrorMessage{Err: err, StatusCode: 500}
	}
	customer.Password = tempPassword

	resp, custErr := this.CustomerRepository.SignUp(customer)
	if custErr.Err != nil {
		return "", custErr
	}

	token, err := this.generateJWTToken(resp)
	if err != nil {
		return "", models.ErrorMessage{Err: err, StatusCode: 500}
	}

	return token, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *customerService) SignIn(customer models.SignInRequest) (string, models.ErrorMessage) {
	resp, custErr := this.CustomerRepository.SignIn(customer)
	if custErr.Err != nil {
		return "", custErr
	}

	if resp.ID == 0 {
		return "", models.ErrorMessage{Err: errors.New("user doesn't exists"), StatusCode: 404}
	}

	isValidPassword := this.verifyPassword(resp.Password, customer.Password)
	if !isValidPassword {
		return "", models.ErrorMessage{Err: errors.New("invalid email or password"), StatusCode: 401}
	}

	token, err := this.generateJWTToken(resp)
	if err != nil {
		return "", models.ErrorMessage{Err: err, StatusCode: 500}
	}
	return token, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *customerService) IsExists(id int) bool {
	return this.CustomerRepository.IsExists("id", strconv.Itoa(id))
}
