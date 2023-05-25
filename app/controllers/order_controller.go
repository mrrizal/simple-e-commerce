package controllers

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
	"e-commerce-api/app/validators"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	orderService      services.OrderService
	productService    services.ProductService
	customerValidator validators.CustomerValidator
	productValidator  validators.ProductValidator
}

func NewOrderController(orderService services.OrderService, productService services.ProductService,
	customerValidator validators.CustomerValidator, productValidator validators.ProductValidator) OrderController {
	return OrderController{
		orderService:      orderService,
		productService:    productService,
		customerValidator: customerValidator,
		productValidator:  productValidator,
	}
}

func (this *OrderController) CreateOrder(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]

	// validate customer
	customerID, custErr := this.customerValidator.ValidateCustomer(token)
	if custErr.Err != nil {
		return models.ErrorResponse(c, custErr)
	}

	// parse body request
	var order models.OrderRequest
	if err := c.BodyParser(&order); err != nil {
		return err
	}
	order.CustomerID = int(customerID)

	// validate product
	temp, custErr := this.productValidator.ValidateProducts(order.ProductsID)
	if custErr.Err != nil {
		return models.ErrorResponse(c, custErr)
	}
	order.ProductsID = temp

	// create order
	orderResp, custErr := this.orderService.CreateOrder(order)
	if custErr.Err != nil {
		return models.ErrorResponse(c, custErr)
	}

	return models.SuccessResponse(c, models.SuccessMessage{Message: orderResp, StatusCode: 201})
}

func (this *OrderController) getProductIds(productData [][]int) ([]int, map[int][]int) {
	productIdsKey := make(map[int]bool)
	orderProductsInfo := make(map[int][]int)
	for _, value := range productData {
		productIdsKey[value[1]] = true
		temp := orderProductsInfo[value[0]]
		temp = append(temp, value[1])
		orderProductsInfo[value[0]] = temp
	}

	result := []int{}
	for key, _ := range productIdsKey {
		result = append(result, key)
	}
	return result, orderProductsInfo
}

func (this *OrderController) mergeProductData(orderData map[int]models.OrderData,
	productData [][]int) models.ErrorMessage {
	productIDs, orderProductsInfo := this.getProductIds(productData)

	tempProducts, err := this.productService.GetMultiple(productIDs)
	if err.Err != nil {
		return err
	}

	products := make(map[int]models.ProductResp)
	for _, product := range tempProducts.Results {
		products[product.ID] = product
	}

	for orderID, productsInfo := range orderProductsInfo {
		temp := orderData[orderID]
		for _, product := range productsInfo {
			temp.Products = append(temp.Products, products[product])
		}
		orderData[orderID] = temp
	}

	return models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *OrderController) Get(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]

	// validate customer
	customerID, err := this.customerValidator.ValidateCustomer(token)
	if err.Err != nil {
		return models.ErrorResponse(c, err)
	}

	orderData, productData, err := this.orderService.GetOrder(customerID)
	if err.Err != nil {
		return models.ErrorResponse(c, err)
	}

	err = this.mergeProductData(orderData, productData)
	if err.Err != nil {
		return models.ErrorResponse(c, err)
	}

	var result models.OrderDataResp
	for _, value := range orderData {
		result.Results = append(result.Results, value)
	}

	return models.SuccessResponse(c, models.SuccessMessage{Message: result, StatusCode: 200})
}
