package repositories

import (
	"context"
	"fmt"

	"e-commerce-api/app/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	CreateOrder(models.OrderRequest) (models.OrderResp, error, int)
	CreateOrderDetail(int, models.OrderRequest) (error, int)
	GetOrder(customerID int) ([]models.Order, error, int)
	GetDetailOrder([]int) ([]models.OrderDetail, error, int)
}

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{db: db}
}

func (this *orderRepository) CreateOrder(order models.OrderRequest) (models.OrderResp, error, int) {
	sqlStmt := `INSERT INTO "order" (status, customer_id) VALUES ($1, $2) RETURNING id`

	var orderID int
	err := this.db.QueryRow(context.Background(), sqlStmt, order.Status, order.CustomerID).Scan(&orderID)
	if err != nil {
		return models.OrderResp{}, err, 500
	}
	return models.OrderResp{ID: orderID, Status: order.Status}, nil, 0
}

func (this *orderRepository) CreateOrderDetail(orderID int, order models.OrderRequest) (error, int) {
	records := [][]interface{}{}
	for _, productID := range order.ProductsID {
		records = append(records, []interface{}{orderID, productID})
	}

	columns := []string{"order_id", "product_id"}

	tx, err := this.db.Begin(context.Background())
	if err != nil {
		return err, 500
	}
	defer tx.Rollback(context.Background())

	copyFrom := pgx.CopyFromRows(records)

	_, err = tx.CopyFrom(context.Background(), pgx.Identifier{"order_detail"}, columns, copyFrom)
	if err != nil {
		return err, 500
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err, 500
	}

	return nil, 0
}

func (this *orderRepository) GetOrder(customerID int) ([]models.Order, error, int) {
	sqlStmt := fmt.Sprintf(`SELECT id, date, customer_id, status FROM "order" WHERE customer_id = %d`, customerID)

	rows, err := this.db.Query(context.Background(), sqlStmt)
	defer rows.Close()
	if err != nil {
		return []models.Order{}, err, 500
	}

	var results []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Date, &order.CustomerID, &order.Status)
		if err != nil {
			return []models.Order{}, err, 500
		}
		results = append(results, order)
	}
	return results, nil, 0
}

func (this *orderRepository) GetDetailOrder(orderIds []int) ([]models.OrderDetail, error, int) {
	sqlStmt := `SELECT id, order_id, product_id FROM order_detail WHERE order_id = ANY($1::integer[])`

	rows, err := this.db.Query(context.Background(), sqlStmt, orderIds)
	defer rows.Close()
	if err != nil {
		return []models.OrderDetail{}, err, 500
	}

	var results []models.OrderDetail
	for rows.Next() {
		var orderDetail models.OrderDetail
		err := rows.Scan(&orderDetail.ID, &orderDetail.OrderID, &orderDetail.ProductID)
		if err != nil {
			return []models.OrderDetail{}, err, 500
		}
		results = append(results, orderDetail)
	}
	return results, nil, 0
}
