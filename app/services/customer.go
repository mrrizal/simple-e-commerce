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

type CustomerService interface {
	SignUp(models.Customer) models.Result
	SignIn(models.SignInRequest) models.Result
	IsExists(int) bool
}

type customerService struct {
	CustomerRepository repositories.CustomerRepository
	config             configs.Config
}

func NewCustomerService(db *pgxpool.Pool) CustomerService {
	return &customerService{
		CustomerRepository: repositories.NewCustomerRepository(db),
		config:             configs.LoadConfig(),
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

func (this *customerService) SignUp(customer models.Customer) models.Result {
	emailIsExists := this.CustomerRepository.IsExists("email", customer.Email)
	if emailIsExists {
		return models.Result{
			Data:       "",
			Err:        errors.New("email already exists"),
			StatusCode: 400,
		}
	}

	tempPassword, err := this.generatePassword(customer.Password)
	if err != nil {
		return models.Result{
			Data:       "",
			Err:        err,
			StatusCode: 500,
		}
	}
	customer.Password = tempPassword

	resp := this.CustomerRepository.SignUp(customer)
	if resp.Err != nil {
		return models.Result{
			Data:       "",
			Err:        resp.Err,
			StatusCode: resp.StatusCode,
		}
	}

	token, err := this.generateJWTToken(resp.Data.(models.Customer))
	if err != nil {
		return models.Result{
			Data:       "",
			Err:        err,
			StatusCode: 500,
		}
	}

	return models.Result{
		Data:       token,
		Err:        nil,
		StatusCode: 0,
	}
}

func (this *customerService) SignIn(customer models.SignInRequest) models.Result {
	resp := this.CustomerRepository.SignIn(customer)
	if resp.Data.(models.Customer).ID == 0 {
		return models.Result{Data: "", Err: errors.New("user doesn't exists"), StatusCode: 404}
	}

	isValidPassword := this.verifyPassword(resp.Data.(models.Customer).Password, customer.Password)
	if !isValidPassword {
		return models.Result{Data: "", Err: errors.New("invalid email or password"), StatusCode: 401}
	}

	token, err := this.generateJWTToken(resp.Data.(models.Customer))
	if err != nil {
		return models.Result{Data: "", Err: err, StatusCode: 500}
	}
	return models.Result{Data: token, Err: nil, StatusCode: 0}
}

func (this *customerService) IsExists(id int) bool {
	return this.CustomerRepository.IsExists("id", strconv.Itoa(id))
}
