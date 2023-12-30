package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"shop/models"
)

type orderProduct struct {
	DB *sql.DB
}

func NewOrdersProduct(db *sql.DB) orderProduct {
	return orderProduct{DB: db}
}

func (o orderProduct) Insert(op models.OrderProducts) (string, error) {
	id := uuid.New()
	if _, err := o.DB.Exec(`insert into order_products(id,order_id,product_id,quantity,price) values($1, $2, $3, $4, $5)`,
		&id, &op.OrderId, &op.ProductId, &op.Quantity, &op.Price); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (o orderProduct) GetByID(id string) (models.OrderProducts, error) {
	orderP := models.OrderProducts{}
	if err := o.DB.QueryRow(`select * from order_products where id = $1`, id).Scan(&orderP.ID, &orderP.OrderId, &orderP.ProductId, &orderP.Quantity, &orderP.Price); err != nil {
		return models.OrderProducts{}, err
	}
	return orderP, nil
}

func (o orderProduct) GetList() ([]models.OrderProducts, error) {
	orders := []models.OrderProducts{}

	rows, err := o.DB.Query(`select * from order_products`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		orderP := models.OrderProducts{}
		if err = rows.Scan(&orderP.ID, &orderP.OrderId, &orderP.ProductId, &orderP.Quantity, &orderP.Price); err != nil {
			return nil, err
		}
		orders = append(orders, orderP)
	}

	return orders, nil
}

func (o orderProduct) Update(orderP models.OrderProducts) error {
	if _, err := o.DB.Exec(`update order_products set order_id = $1, product_id = $2, quantity = $3, price = $4 where id = $5`,
		&orderP.OrderId, &orderP.ProductId, &orderP.Quantity, &orderP.Price, &orderP.ID); err != nil {
		return err
	}
	return nil
}

func (o orderProduct) Delete(id string) error {
	if _, err := o.DB.Exec(`delete from order_products where id = $1`, id); err != nil {
		return err
	}
	return nil
}
