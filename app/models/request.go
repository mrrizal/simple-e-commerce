package models

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OrderRequest struct {
	Status     string `json:"status"`
	CustomerID int    `json:"customer_id"`
	ProductsID []int  `json:"products_id"`
}
