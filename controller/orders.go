package controller

import (
	"fmt"
	"log"
	"shop/models"
)

func (c Controller) CreateOrders() {
	order := getOrderInfo()

	if _, err := c.Store.OrdersRepo.Insert(order); err != nil {
		log.Fatalln("error is while inserting", err.Error())
		return
	}
	fmt.Println("user added")
}

func (c Controller) GetOrdersByID() {
	id := ""
	fmt.Print("input id: ")
	fmt.Scan(&id)

	order, err := c.Store.OrdersRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while get by id", err)
		return
	}
	fmt.Println("order: ", order)
}

func (c Controller) GetOrdersList() {
	orders, err := c.Store.OrdersRepo.GetList()
	if err != nil {
		fmt.Print("error is while get list", err.Error())
		return
	}
	fmt.Println("orders: ", orders)
}

func (c Controller) UpdateOrder() {
	order := getOrderInfo()

	if err := c.Store.OrdersRepo.Update(order); err != nil {
		fmt.Print("error is while updating", err.Error())
		return
	}
	fmt.Println("order updated")
}

func (c Controller) DeleteOrder() {
	id := ""
	fmt.Print("input id: ")
	fmt.Scan(&id)

	if err := c.Store.OrdersRepo.Delete(id); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return
	}
	fmt.Println("user deleted!")
}

func getOrderInfo() models.Order {
	var (
		id, user_id string
		cmd, amount int
	)
	fmt.Print(`
					1 - create
					2 - update
`)
	fmt.Scan(&cmd)

	if cmd == 2 {
		fmt.Print("input id: ")
		fmt.Scan(&id)

		fmt.Print("input amount: ")
		fmt.Scan(&amount)

		fmt.Print("input user id: ")
		fmt.Scan(&user_id)
	} else if cmd == 1 {
		fmt.Print("input amount: ")
		fmt.Scan(&amount)

		fmt.Print("input user id: ")
		fmt.Scan(&user_id)
	} else {
		fmt.Println(cmd, "not found")
	}

	if id != "" {
		return models.Order{
			ID:     id,
			Amount: amount,
			UserID: user_id,
		}
	}
	return models.Order{
		Amount: amount,
		UserID: user_id,
	}

}
