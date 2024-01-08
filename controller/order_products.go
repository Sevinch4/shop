package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shop/models"
)

func (c Controller) OrderProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateOrderProduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			c.GetOrderProductsList(w, r)
		} else {
			c.GetOrderProductByID(w, r)
		}
	case http.MethodPut:
		c.UpdateOrderProduct(w, r)
	case http.MethodDelete:
		c.DeleteOrderProduct(w, r)
	}
}

func (c Controller) CreateOrderProduct(w http.ResponseWriter, r *http.Request) {
	order := models.OrderProducts{}

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.OrderProduct.Insert(order)
	if err != nil {
		log.Fatalln("error is while inserting", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := c.Store.OrderProduct.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c Controller) GetOrderProductByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	order, err := c.Store.OrderProduct.GetByID(id)
	if err != nil {
		fmt.Println("error is while get by id", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, order)
}

func (c Controller) GetOrderProductsList(w http.ResponseWriter, r *http.Request) {
	orders, err := c.Store.OrderProduct.GetList()
	if err != nil {
		fmt.Print("error is while get list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, orders)
}

func (c Controller) UpdateOrderProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	orderProduct := models.OrderProducts{}

	if err := json.NewDecoder(r.Body).Decode(&orderProduct); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if orderProduct.ID != id {
		fmt.Println("car ID not mismatch")
		handleResponse(w, http.StatusBadRequest, orderProduct.ID)
		return
	}

	if err := c.Store.OrderProduct.Update(orderProduct); err != nil {
		fmt.Print("error is while updating", err.Error())
		return
	}

	resp, err := c.Store.OrderProduct.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c Controller) DeleteOrderProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Store.OrderProduct.Delete(id); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return
	}

	handleResponse(w, http.StatusOK, id)
}
