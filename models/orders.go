package models

import "time"

type Order struct {
	ID        string
	Amount    int
	UserID    string
	CreatedAt time.Time
}
