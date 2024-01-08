package models

type Product struct {
	ID    string `json:"id"`
	Price int    `json:"price"`
	Name  string `json:"name"`
}
