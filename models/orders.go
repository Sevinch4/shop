package models

import "time"

type Order struct {
	ID        string    `json:"id"`
	Amount    int       `json:"amount"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderProducts struct {
	ID        string `json:"id"`
	OrderId   string `json:"order_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
}
