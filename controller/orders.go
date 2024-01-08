package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shop/models"
)

func (c Controller) Orders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateOrders(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			c.GetOrdersList(w, r)
		} else {
			c.GetOrdersByID(w, r)
		}
	case http.MethodPut:
		c.UpdateOrder(w, r)
	case http.MethodDelete:
		c.DeleteOrder(w, r)
	}
}

func (c Controller) CreateOrders(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.OrdersRepo.Insert(order)
	if err != nil {
		log.Fatalln("error is while inserting", err.Error())
		return
	}

	resp, err := c.Store.OrdersRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c Controller) GetOrdersByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	order, err := c.Store.OrdersRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while get by id", err)
		return
	}

	handleResponse(w, http.StatusOK, order)
}

func (c Controller) GetOrdersList(w http.ResponseWriter, r *http.Request) {
	orders, err := c.Store.OrdersRepo.GetList()
	if err != nil {
		fmt.Print("error is while get list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, orders)
}

func (c Controller) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	order := models.Order{}

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if order.ID != id {
		fmt.Println("car ID not mismatch")
		handleResponse(w, http.StatusBadRequest, order.ID)
		return
	}

	if err := c.Store.OrdersRepo.Update(order); err != nil {
		fmt.Print("error is while updating", err.Error())
		return
	}

	resp, err := c.Store.OrdersRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is ehile getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c Controller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Store.OrdersRepo.Delete(id); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return
	}

	handleResponse(w, http.StatusOK, id)
}
