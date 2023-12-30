package controller

import (
	"fmt"
	"shop/models"
)

func (c Controller) CreateProduct() {
	product := getProductInfo()

	if !checkPrice(product.Price) {
		fmt.Println("error is while checking price, price format not correct")
		return
	}

	if _, err := c.Store.ProductsRepo.Insert(product); err != nil {
		fmt.Println("error is while inserting product", err.Error())
		return
	}

	fmt.Println("product added")
}

func (c Controller) GetProductByID() {
	id := ""
	fmt.Print("input id: ")
	fmt.Scan(&id)

	product, err := c.Store.ProductsRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}
	fmt.Println("product: ", product)
}

func (c Controller) GetProductList() {
	products, err := c.Store.ProductsRepo.GetList()
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		return
	}
	fmt.Println(products)
}

func (c Controller) UpdateProduct() {
	prod := getProductInfo()
	if err := c.Store.ProductsRepo.Update(prod); err != nil {
		fmt.Println("error is while updating data", err.Error())
		return
	}
	fmt.Println("product updated")
}

func (c Controller) DeleteProduct() {
	id := ""
	fmt.Print("input id: ")
	fmt.Scan(&id)

	if err := c.Store.ProductsRepo.Delete(id); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return
	}
	fmt.Println("user deleted!")
}

func checkPrice(price int) bool {
	if price < 0 {
		return false
	}
	return true
}

func getProductInfo() models.Product {
	var (
		id, name   string
		price, cmd int
	)
	fmt.Print(`
					1 - create
					2 - update
`)
	fmt.Scan(&cmd)

	if cmd == 2 {
		fmt.Print("input id: ")
		fmt.Scan(&id)

		fmt.Print("input product name: ")
		fmt.Scan(&name)

		fmt.Print("input product price: ")
		fmt.Scan(&price)
	} else if cmd == 1 {
		fmt.Print("input product name: ")
		fmt.Scan(&name)

		fmt.Print("input product price: ")
		fmt.Scan(&price)
	} else {
		fmt.Println("not found")
	}

	if id != "" {
		return models.Product{
			ID:    id,
			Price: price,
			Name:  name,
		}
	}

	return models.Product{
		Price: price,
		Name:  name,
	}
}
