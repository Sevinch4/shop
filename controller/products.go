package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shop/models"
	"shop/pkg/check"
)

func (c Controller) Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateProduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			c.GetProductList(w, r)
		} else {
			c.GetProductByID(w, r)
		}
	case http.MethodPut:
		c.UpdateProduct(w, r)
	case http.MethodDelete:
		c.DeleteProduct(w, r)
	}
}

func (c Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !check.Price(product.Price) {
		fmt.Println("error is while checking price, price format not correct")
		handleResponse(w, http.StatusBadRequest, product.Price)
		return
	}

	id, err := c.Store.ProductsRepo.Insert(product)
	if err != nil {
		fmt.Println("error is while inserting product", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := c.Store.ProductsRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}

func (c Controller) GetProductByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	product, err := c.Store.ProductsRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, product)
}

func (c Controller) GetProductList(w http.ResponseWriter, r *http.Request) {
	products, err := c.Store.ProductsRepo.GetList()
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, products)
}

func (c Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	product := models.Product{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if product.ID != id {
		fmt.Println("car ID not mismatch")
		handleResponse(w, http.StatusBadRequest, product.ID)
		return
	}

	if !check.Price(product.Price) {
		fmt.Println("error is product format not correct", product.Price)
		handleResponse(w, http.StatusBadRequest, product.Price)
		return
	}

	if err := c.Store.ProductsRepo.Update(product); err != nil {
		fmt.Println("error is while updating data", err.Error())
		return
	}

	resp, err := c.Store.ProductsRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is ehile getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Store.ProductsRepo.Delete(id); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return
	}

	handleResponse(w, http.StatusOK, id)
}
