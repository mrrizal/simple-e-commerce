package services

import (
	"e-commerce-api/app/configs"
	"e-commerce-api/app/models"
	"e-commerce-api/app/repositories"
	"strconv"

	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	CustomerRepository repositories.CustomerRepository
	config             configs.Config
}

func NewCustomerService(db *pgxpool.Pool) CustomerService {
	return CustomerService{
		CustomerRepository: repositories.NewCustomerRepository(db),
		config:             configs.LoadConfig(),
	}
}

func (this *CustomerService) generatePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (this *CustomerService) verifyPassword(customerPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(customerPassword), []byte(password))
	return err == nil
}

func (this *CustomerService) generateJWTToken(customer models.Customer) (string, error) {
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

func (this *CustomerService) SignUp(customer models.Customer) (string, error, int) {
	emailIsExists := this.CustomerRepository.IsExists("email", customer.Email)
	if emailIsExists {
		return "", errors.New("email already exists"), 400
	}

	tempPassword, err := this.generatePassword(customer.Password)
	if err != nil {
		return "", err, 500
	}
	customer.Password = tempPassword

	customer, err, statusCode := this.CustomerRepository.SignUp(customer)
	if err != nil {
		return "", err, statusCode
	}

	token, err := this.generateJWTToken(customer)
	if err != nil {
		return "", err, 500
	}
	return token, nil, 0
}

func (this *CustomerService) SignIn(customer models.SignInRequest) (string, error, int) {
	cust := this.CustomerRepository.SignIn(customer)
	if cust.ID == 0 {
		return "", errors.New("user doesn't exists"), 404
	}

	isValidPassword := this.verifyPassword(cust.Password, customer.Password)
	if !isValidPassword {
		return "", errors.New("invalid email or password"), 401
	}

	token, err := this.generateJWTToken(cust)
	if err != nil {
		return "", err, 500
	}
	return token, nil, 0
}

func (this *CustomerService) IsExists(id int) bool {
	return this.CustomerRepository.IsExists("id", strconv.Itoa(id))
}
