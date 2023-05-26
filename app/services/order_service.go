// this package contain function that do the bussiness logic
package services

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/repositories"
)

type OrderService interface {
	CreateOrder(models.OrderRequest) (models.OrderResp, models.ErrorMessage)
	GetOrder(int) (map[int]models.OrderData, [][]int, models.ErrorMessage)
}

type orderService struct {
	OrderRepository repositories.OrderRepository
}

func NewOrderService(repository repositories.OrderRepository) OrderService {
	return &orderService{OrderRepository: repository}
}

func (this *orderService) CreateOrder(order models.OrderRequest) (models.OrderResp, models.ErrorMessage) {
	orderResp, err := this.OrderRepository.CreateOrder(order)
	if err.Err != nil {
		return models.OrderResp{}, err
	}

	resp := this.OrderRepository.CreateOrderDetail(orderResp.ID, order)
	if resp.Err != nil {
		return models.OrderResp{}, err
	}

	return orderResp, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *orderService) GetOrder(customerID int) (map[int]models.OrderData, [][]int, models.ErrorMessage) {
	orders, err := this.OrderRepository.GetOrder(customerID)
	if err.Err != nil {
		return make(map[int]models.OrderData), [][]int{}, err
	}
	orderIds := []int{}

	mapOrders := make(map[int]models.OrderData)
	for _, order := range orders {
		var temp models.OrderData
		temp.ID = order.ID
		temp.CustomerID = order.CustomerID
		temp.Date = order.Date
		temp.Status = order.Status
		mapOrders[order.ID] = temp
		orderIds = append(orderIds, order.ID)
	}

	orderDetails, err := this.OrderRepository.GetDetailOrder(orderIds)
	if err.Err != nil {
		return make(map[int]models.OrderData), [][]int{}, err
	}

	productsData := [][]int{}
	for _, orderDetail := range orderDetails {
		productsData = append(productsData, []int{orderDetail.OrderID, orderDetail.ProductID})
	}
	return mapOrders, productsData, models.ErrorMessage{Err: nil, StatusCode: 0}
}
