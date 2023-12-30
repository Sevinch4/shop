package controller

import (
	"fmt"
	"log"
	"shop/models"
)

func (c Controller) CreateOrderProduct() {
	order := getOrderProductInfo()

	if _, err := c.Store.OrderProduct.Insert(order); err != nil {
		log.Fatalln("error is while inserting", err.Error())
		return
	}
	fmt.Println("user added")
}

func (c Controller) GetOrderProductByID() {
	id := ""
	fmt.Print("input id: ")
	fmt.Scan(&id)

	order, err := c.Store.OrderProduct.GetByID(id)
	if err != nil {
		fmt.Println("error is while get by id", err)
		return
	}
	fmt.Println("order: ", order)
}

func (c Controller) GetOrderProductsList() {
	orders, err := c.Store.OrderProduct.GetList()
	if err != nil {
		fmt.Print("error is while get list", err.Error())
		return
	}
	fmt.Println("orders: ", orders)
}

func (c Controller) UpdateOrderProduct() {
	order := getOrderProductInfo()

	if err := c.Store.OrderProduct.Update(order); err != nil {
		fmt.Print("error is while updating", err.Error())
		return
	}
	fmt.Println("order updated")
}

func (c Controller) DeleteOrderProduct() {
	id := ""
	fmt.Print("input id: ")
	fmt.Scan(&id)

	if err := c.Store.OrderProduct.Delete(id); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return
	}
	fmt.Println("user deleted!")
}

func getOrderProductInfo() models.OrderProducts {
	var (
		id, order_id, pro_id string
		cmd, price, quantity int
	)
	fmt.Print(`
					1 - create
					2 - update
`)
	fmt.Scan(&cmd)

	if cmd == 2 {
		fmt.Print("input id: ")
		fmt.Scan(&id)

		fmt.Print("input order id: ")
		fmt.Scan(&order_id)

		fmt.Print("input product id: ")
		fmt.Scan(&pro_id)

		fmt.Print("input quantity: ")
		fmt.Scan(&quantity)

		fmt.Print("input price: ")
		fmt.Scan(&price)
	} else if cmd == 1 {
		fmt.Print("input order id: ")
		fmt.Scan(&order_id)

		fmt.Print("input product id: ")
		fmt.Scan(&pro_id)

		fmt.Print("input quantity: ")
		fmt.Scan(&quantity)

		fmt.Print("input price: ")
		fmt.Scan(&price)
	} else {
		fmt.Println(cmd, "not found")
	}

	if id != "" {
		return models.OrderProducts{
			ID:        id,
			OrderId:   order_id,
			ProductId: pro_id,
			Quantity:  quantity,
			Price:     price,
		}
	}
	return models.OrderProducts{
		OrderId:   order_id,
		ProductId: pro_id,
		Quantity:  quantity,
		Price:     price,
	}

}
