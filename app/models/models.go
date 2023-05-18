package models

import (
	"database/sql"
	"time"
)

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Price       string         `json:"price"`
	Description sql.NullString `json:"description"`
	Image       sql.NullString `json:"image"`
}

type ProductList struct {
	Results []ProductResp `json:"results"`
}

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Products   []int     `json:"products"`
	Status     string    `json:"status"`
	Date       time.Time `json:"date"`
}

type OrderDetail struct {
	ID        int `json:"id"`
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
}

type OrderData struct {
	ID         int           `json:"id"`
	CustomerID int           `json:"customer_id"`
	Status     string        `json:"status"`
	Date       time.Time     `json:"date"`
	Products   []ProductResp `json:"products"`
}
