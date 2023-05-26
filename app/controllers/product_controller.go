package controllers

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/services"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return ProductController{ProductService: productService}
}

func (this *ProductController) Get(productID int) interface{} {
	product, custErr := this.ProductService.Get(productID)
	if custErr.Err != nil {
		return custErr
	}

	return models.SuccessMessage{Message: product, StatusCode: 200}
}

func (this *ProductController) GetAll() interface{} {
	product, err := this.ProductService.GetAll()
	if err.Err != nil {
		return err
	}

	return models.SuccessMessage{Message: product, StatusCode: 200}
}

func (this *ProductController) GetMultiple(ids []int) interface{} {
	product, err := this.ProductService.GetMultiple(ids)
	if err.Err != nil {
		return err
	}

	return models.SuccessMessage{Message: product, StatusCode: 200}
}
