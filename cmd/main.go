package main

import (
	_ "github.com/lib/pq"
	"log"
	"shop/config"
	"shop/controller"
	"shop/storage/postgres"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer store.DB.Close()

	con := controller.New(store)

	//con.CreateUser()
	//con.GetUserByID()
	//con.GetUserList()
	//con.UpdateUser()
	//con.DeleteUser()

	//con.CreateOrders()
	//con.GetOrdersByID()
	//con.GetOrdersList()
	con.UpdateOrder()
	//con.DeleteOrder()

	//con.CreateProduct()
	//con.GetProductByID()
	//con.GetProductList()
	//con.UpdateProduct()-----error
	//con.DeleteProduct()

	//con.CreateOrderProduct()
	//con.GetOrderProductByID()
	//con.GetOrderProductsList()
	//con.UpdateOrderProduct()
	//con.DeleteOrderProduct()

}
