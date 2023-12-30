package postgres

import (
	"database/sql"
	"fmt"
	"shop/config"
)

type Store struct {
	DB           *sql.DB
	UserRepo     usersRepo
	OrdersRepo   ordersRepo
	ProductsRepo productsRepo
	OrderProduct orderProduct
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host = %s port = %s user = %s password = %s database = %s sslmode = disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	user := NewUsersRepo(db)
	order := NewOrdersRepo(db)
	product := NewProductsRepo(db)
	ordersProduct := NewOrdersProduct(db)

	return Store{DB: db,
		UserRepo:     user,
		OrdersRepo:   order,
		ProductsRepo: product,
		OrderProduct: ordersProduct,
	}, err
}
