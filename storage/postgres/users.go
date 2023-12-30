package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"shop/models"
)

type usersRepo struct {
	DB *sql.DB
}

func NewUsersRepo(db *sql.DB) usersRepo {
	return usersRepo{DB: db}
}

func (u usersRepo) Insert(us models.User) (string, error) {
	id := uuid.New()
	if _, err := u.DB.Exec(`insert into users (id,first_name,last_name,email,phone) values ($1,$2,$3,$4,$5)`,
		&id, &us.FirstName, &us.LastName, &us.Email, &us.Phone); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (u usersRepo) GetByID(id string) (models.User, error) {
	us := models.User{}
	if err := u.DB.QueryRow(`select * from users where id = $1`, id).Scan(&us.ID, &us.FirstName, &us.LastName, &us.Email, &us.Phone); err != nil {
		return models.User{}, err
	}
	return us, nil
}

func (u usersRepo) GetList() ([]models.User, error) {
	users := []models.User{}

	rows, err := u.DB.Query(`select * from users`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		us := models.User{}
		if err := rows.Scan(&us.ID, &us.FirstName, &us.LastName, &us.Email, &us.Phone); err != nil {
			return nil, err
		}
		users = append(users, us)
	}
	return users, nil
}

func (u usersRepo) Update(us models.User) error {
	if _, err := u.DB.Exec(`update users set first_name = $1,last_name = $2, email = $3, phone = $4 where id = $5`,
		&us.FirstName, &us.LastName, &us.Email, &us.Phone, &us.ID); err != nil {
		return err
	}
	return nil
}

func (u usersRepo) Delete(id string) error {
	if _, err := u.DB.Exec(`delete from users where id = $1`, id); err != nil {
		return err
	}
	return nil
}
