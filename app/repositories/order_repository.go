package repositories

import (
	"context"
	"fmt"

	"e-commerce-api/app/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	CreateOrder(models.OrderRequest) (models.OrderResp, models.CustomError)
	CreateOrderDetail(int, models.OrderRequest) models.CustomError
	GetOrder(customerID int) ([]models.Order, models.CustomError)
	GetDetailOrder([]int) ([]models.OrderDetail, models.CustomError)
}

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{db: db}
}

func (this *orderRepository) CreateOrder(order models.OrderRequest) (models.OrderResp, models.CustomError) {
	sqlStmt := `INSERT INTO "order" (status, customer_id) VALUES ($1, $2) RETURNING id`

	var orderID int
	err := this.db.QueryRow(context.Background(), sqlStmt, order.Status, order.CustomerID).Scan(&orderID)
	if err != nil {
		return models.OrderResp{}, models.CustomError{Err: err, StatusCode: 500}
	}
	return models.OrderResp{ID: orderID, Status: order.Status}, models.CustomError{Err: nil, StatusCode: 0}
}

func (this *orderRepository) CreateOrderDetail(orderID int, order models.OrderRequest) models.CustomError {
	records := [][]interface{}{}
	for _, productID := range order.ProductsID {
		records = append(records, []interface{}{orderID, productID})
	}

	columns := []string{"order_id", "product_id"}

	tx, err := this.db.Begin(context.Background())
	if err != nil {
		return models.CustomError{Err: err, StatusCode: 500}
	}
	defer tx.Rollback(context.Background())

	copyFrom := pgx.CopyFromRows(records)

	_, err = tx.CopyFrom(context.Background(), pgx.Identifier{"order_detail"}, columns, copyFrom)
	if err != nil {
		return models.CustomError{Err: err, StatusCode: 500}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return models.CustomError{Err: err, StatusCode: 500}
	}

	return models.CustomError{Err: nil, StatusCode: 0}
}

func (this *orderRepository) GetOrder(customerID int) ([]models.Order, models.CustomError) {
	sqlStmt := fmt.Sprintf(`SELECT id, date, customer_id, status FROM "order" WHERE customer_id = %d`, customerID)

	rows, err := this.db.Query(context.Background(), sqlStmt)
	defer rows.Close()
	if err != nil {
		return []models.Order{}, models.CustomError{Err: err, StatusCode: 500}
	}

	var results []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Date, &order.CustomerID, &order.Status)
		if err != nil {
			return []models.Order{}, models.CustomError{Err: err, StatusCode: 500}

		}
		results = append(results, order)
	}
	return results, models.CustomError{Err: nil, StatusCode: 0}
}

func (this *orderRepository) GetDetailOrder(orderIds []int) ([]models.OrderDetail, models.CustomError) {
	sqlStmt := `SELECT id, order_id, product_id FROM order_detail WHERE order_id = ANY($1::integer[])`

	rows, err := this.db.Query(context.Background(), sqlStmt, orderIds)
	defer rows.Close()
	if err != nil {
		return []models.OrderDetail{}, models.CustomError{Err: err, StatusCode: 500}
	}

	var results []models.OrderDetail
	for rows.Next() {
		var orderDetail models.OrderDetail
		err := rows.Scan(&orderDetail.ID, &orderDetail.OrderID, &orderDetail.ProductID)
		if err != nil {
			return []models.OrderDetail{}, models.CustomError{Err: err, StatusCode: 500}
		}
		results = append(results, orderDetail)
	}
	return results, models.CustomError{Err: nil, StatusCode: 0}
}
