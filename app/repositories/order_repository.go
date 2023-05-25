package repositories

import (
	"context"
	"fmt"

	"e-commerce-api/app/database"
	"e-commerce-api/app/models"
)

type OrderRepository interface {
	CreateOrder(models.OrderRequest) (models.OrderResp, models.ErrorMessage)
	CreateOrderDetail(int, models.OrderRequest) models.ErrorMessage
	GetOrder(customerID int) ([]models.Order, models.ErrorMessage)
	GetDetailOrder([]int) ([]models.OrderDetail, models.ErrorMessage)
}

type orderRepository struct {
	db database.DB
}

func NewOrderRepository(db database.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (this *orderRepository) CreateOrder(order models.OrderRequest) (models.OrderResp, models.ErrorMessage) {
	sqlStmt := `INSERT INTO "order" (status, customer_id) VALUES ($1, $2) RETURNING id`

	var orderID int
	err := this.db.QueryRow(context.Background(), sqlStmt, order.Status, order.CustomerID).Scan(&orderID)
	if err != nil {
		return models.OrderResp{}, models.ErrorMessage{Err: err, StatusCode: 500}
	}
	return models.OrderResp{ID: orderID, Status: order.Status}, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *orderRepository) CreateOrderDetail(orderID int, order models.OrderRequest) models.ErrorMessage {
	records := [][]interface{}{}
	for _, productID := range order.ProductsID {
		records = append(records, []interface{}{orderID, productID})
	}

	columns := []string{"order_id", "product_id"}

	tx, err := this.db.Begin(context.Background())
	if err != nil {
		defer tx.Rollback(context.Background())
		return models.ErrorMessage{Err: err, StatusCode: 500}
	}

	_, err = tx.BulkInsert(context.Background(), "order_detail", columns, records)
	if err != nil {
		defer tx.Rollback(context.Background())
		return models.ErrorMessage{Err: err, StatusCode: 500}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		defer tx.Rollback(context.Background())
		return models.ErrorMessage{Err: err, StatusCode: 500}
	}

	return models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *orderRepository) GetOrder(customerID int) ([]models.Order, models.ErrorMessage) {
	sqlStmt := fmt.Sprintf(`SELECT id, date, customer_id, status FROM "order" WHERE customer_id = %d`, customerID)

	rows, err := this.db.Query(context.Background(), sqlStmt)
	defer rows.Close()
	if err != nil {
		return []models.Order{}, models.ErrorMessage{Err: err, StatusCode: 500}
	}

	var results []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Date, &order.CustomerID, &order.Status)
		if err != nil {
			return []models.Order{}, models.ErrorMessage{Err: err, StatusCode: 500}

		}
		results = append(results, order)
	}
	return results, models.ErrorMessage{Err: nil, StatusCode: 0}
}

func (this *orderRepository) GetDetailOrder(orderIds []int) ([]models.OrderDetail, models.ErrorMessage) {
	sqlStmt := `SELECT id, order_id, product_id FROM order_detail WHERE order_id = ANY($1::integer[])`

	rows, err := this.db.Query(context.Background(), sqlStmt, orderIds)
	defer rows.Close()
	if err != nil {
		return []models.OrderDetail{}, models.ErrorMessage{Err: err, StatusCode: 500}
	}

	var results []models.OrderDetail
	for rows.Next() {
		var orderDetail models.OrderDetail
		err := rows.Scan(&orderDetail.ID, &orderDetail.OrderID, &orderDetail.ProductID)
		if err != nil {
			return []models.OrderDetail{}, models.ErrorMessage{Err: err, StatusCode: 500}
		}
		results = append(results, orderDetail)
	}
	return results, models.ErrorMessage{Err: nil, StatusCode: 0}
}
