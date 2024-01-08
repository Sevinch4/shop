package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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

	http.HandleFunc("/user", con.User)
	http.HandleFunc("/product", con.Product)
	http.HandleFunc("/orders", con.Orders)
	http.HandleFunc("/ordered_products", con.OrderProducts)

	fmt.Println("server running......")
	http.ListenAndServe("localhost:8080", nil)

}
