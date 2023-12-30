package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"shop/models"
)

type ordersRepo struct {
	DB *sql.DB
}

func NewOrdersRepo(db *sql.DB) ordersRepo {
	return ordersRepo{DB: db}
}

func (o ordersRepo) Insert(order models.Order) (string, error) {
	id := uuid.New()
	if _, err := o.DB.Exec(`insert into orders(id,amount,user_id) values($1,$2,$3)`,
		&id, &order.Amount, &order.UserID); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (o ordersRepo) GetByID(id string) (models.Order, error) {
	order := models.Order{}
	if err := o.DB.QueryRow(`select * from orders where id = $1 `, id).Scan(&order.ID, &order.Amount, &order.UserID, &order.CreatedAt); err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (o ordersRepo) GetList() ([]models.Order, error) {
	orders := []models.Order{}

	rows, err := o.DB.Query(`select * from orders`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		us := models.Order{}
		if err := rows.Scan(&us.ID, &us.Amount, &us.UserID, &us.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, us)
	}
	return orders, nil
}

func (o ordersRepo) Update(order models.Order) error {
	if _, err := o.DB.Exec(`update orders set amount = $1, user_id = $2, created_at = $3 where id = $4`, &order.ID, &order.Amount, &order.UserID, &order.CreatedAt); err != nil {
		return err
	}
	return nil
}
func (o ordersRepo) Delete(id string) error {
	if _, err := o.DB.Exec(`delete from orders where id = $1`, id); err != nil {
		return err
	}
	return nil
}
