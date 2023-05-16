package services

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderService struct {
	OrderRepository repositories.OrderRepository
}

func NewOrderService(db *pgxpool.Pool) OrderService {
	return OrderService{
		OrderRepository: repositories.NewOrderRepository(db),
	}
}

func (this *OrderService) CreateOrder(order models.OrderRequest) (models.OrderResp, error, int) {
	orderResp, err, statusCode := this.OrderRepository.CreateOrder(order)
	if err != nil {
		return models.OrderResp{}, err, statusCode
	}

	err, statusCode = this.OrderRepository.CreateOrderDetail(orderResp.ID, order)
	if err != nil {
		return models.OrderResp{}, err, statusCode
	}

	return orderResp, nil, 0
}

func (this *OrderService) GetOrder(customerID int) (map[int]models.OrderData, [][]int, error, int) {
	orders, err, statusCode := this.OrderRepository.GetOrder(customerID)
	if err != nil {
		return make(map[int]models.OrderData), [][]int{}, err, statusCode
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

	orderDetails, err, statusCode := this.OrderRepository.GetDetailOrder(orderIds)
	if err != nil {
		return make(map[int]models.OrderData), [][]int{}, err, statusCode
	}

	productsData := [][]int{}
	for _, orderDetail := range orderDetails {
		productsData = append(productsData, []int{orderDetail.OrderID, orderDetail.ProductID})
	}
	return mapOrders, productsData, nil, 0
}
