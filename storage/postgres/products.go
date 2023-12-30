package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"shop/models"
)

type productsRepo struct {
	DB *sql.DB
}

func NewProductsRepo(db *sql.DB) productsRepo {
	return productsRepo{
		DB: db,
	}
}

func (p productsRepo) Insert(prod models.Product) (string, error) {
	id := uuid.New()
	if _, err := p.DB.Exec(`insert into products (id,price,name) values($1, $2, $3)`,
		&id, &prod.Price, &prod.Name); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (p productsRepo) GetByID(id string) (models.Product, error) {
	prod := models.Product{}

	if err := p.DB.QueryRow(`select * from products where id = $1`, id).Scan(&prod.ID, &prod.Price, &prod.Name); err != nil {
		return models.Product{}, err
	}
	return prod, nil
}

func (p productsRepo) GetList() ([]models.Product, error) {
	products := []models.Product{}

	rows, err := p.DB.Query(`select * from products`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		prod := models.Product{}
		if err := rows.Scan(&prod.ID, &prod.Price, &prod.Name); err != nil {
			return nil, err
		}
		products = append(products, prod)
	}
	return products, nil
}

func (p productsRepo) Update(prod models.Product) error {
	if _, err := p.DB.Exec(`update products set price = $1,name = $2 where id = $3`,
		&prod.ID, &prod.Price, &prod.Name); err != nil {
		return err
	}
	return nil
}

func (p productsRepo) Delete(id string) error {
	if _, err := p.DB.Exec(`delete from products where id = $1`, id); err != nil {
		return err
	}
	return nil
}
