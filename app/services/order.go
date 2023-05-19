package services

import (
	"e-commerce-api/app/models"
	"e-commerce-api/app/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderService interface {
	CreateOrder(models.OrderRequest) (models.OrderResp, models.CustomError)
	GetOrder(int) (map[int]models.OrderData, [][]int, models.CustomError)
}

type orderService struct {
	OrderRepository repositories.OrderRepository
}

func NewOrderService(db *pgxpool.Pool) OrderService {
	return &orderService{
		OrderRepository: repositories.NewOrderRepository(db),
	}
}

func (this *orderService) CreateOrder(order models.OrderRequest) (models.OrderResp, models.CustomError) {
	orderResp, err := this.OrderRepository.CreateOrder(order)
	if err.Err != nil {
		return models.OrderResp{}, err
	}

	resp := this.OrderRepository.CreateOrderDetail(orderResp.ID, order)
	if resp.Err != nil {
		return models.OrderResp{}, err
	}

	return orderResp, models.CustomError{Err: nil, StatusCode: 0}
}

func (this *orderService) GetOrder(customerID int) (map[int]models.OrderData, [][]int, models.CustomError) {
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
	return mapOrders, productsData, models.CustomError{Err: nil, StatusCode: 0}
}
