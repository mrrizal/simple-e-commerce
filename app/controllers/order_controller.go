package controllers

import (
	"e-commerce-api/app/configs"
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
	"e-commerce-api/app/validators"
	"e-commerce-api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	OrderService      services.OrderService
	ProductService    services.ProductService
	CustomerValidator validators.CustomerValidator
	ProductValidator  validators.ProductValidator
	Config            configs.Config
}

func (this *OrderController) CreateOrder(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]

	// validate customer
	customerID, custErr := this.CustomerValidator.ValidateCustomer(token)
	if custErr.Err != nil {
		errResp := models.ErrorResponse{
			Message:    custErr.Err.Error(),
			StatusCode: custErr.StatusCode,
		}
		return errResp.Resp(c)
	}

	// parse body request
	var order models.OrderRequest
	if err := c.BodyParser(&order); err != nil {
		return err
	}
	order.CustomerID = int(customerID)

	// validate product
	temp, err, statusCode := this.ProductValidator.ValidateProducts(order.ProductsID)
	if err != nil {
		return utils.ErrorResp(c, err.Error(), statusCode)
	}
	order.ProductsID = temp

	// create order
	orderResp, custErr := this.OrderService.CreateOrder(order)
	if custErr.Err != nil {
		errorResp := models.ErrorResponse{
			Message:    custErr.Err.Error(),
			StatusCode: custErr.StatusCode,
		}
		return errorResp.Resp(c)
	}

	successResp := models.SuccessResponse{
		Message:    orderResp,
		StatusCode: 201,
	}
	return successResp.Resp(c)
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

func (this *OrderController) mergeProductData(orderData map[int]models.OrderData, productData [][]int) (error, int) {
	productIDs, orderProductsInfo := this.getProductIds(productData)

	tempProducts, err, statusCode := this.ProductService.GetMultiple(productIDs)
	if err != nil {
		return err, statusCode
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

	return nil, 0
}

func (this *OrderController) Get(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]

	// validate customer
	customerID, err, statusCode := this.CustomerValidator.ValidateCustomer(token)
	if err != nil {
		return utils.ErrorResp(c, err.Error(), statusCode)
	}

	orderData, productData, custErr := this.OrderService.GetOrder(customerID)
	if err != nil {
		errResp := models.ErrorResponse{
			Message:    custErr.Err.Error(),
			StatusCode: custErr.StatusCode,
		}
		return errResp.Resp(c)
	}

	this.mergeProductData(orderData, productData)
	var result models.OrderDataResp
	for _, value := range orderData {
		result.Results = append(result.Results, value)
	}

	successResp := models.SuccessResponse{
		Message:    result,
		StatusCode: 200,
	}
	return successResp.Resp(c)
}
