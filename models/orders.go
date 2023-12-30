package models

import "time"

type Order struct {
	ID        string
	Amount    int
	UserID    string
	CreatedAt time.Time
}

type OrderProducts struct {
	ID        string
	OrderId   string
	ProductId string
	Quantity  int
	Price     int
}
