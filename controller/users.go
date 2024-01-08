package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shop/models"
	"shop/pkg/check"
)

func (c Controller) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateUser(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			c.GetUserList(w, r)
		} else {
			c.GetUserByID(w, r)
		}
	case http.MethodPut:
		c.UpdateUser(w, r)
	case http.MethodDelete:
		c.DeleteUser(w, r)
	}
}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("error is while decoding data", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !check.PhoneNumber(user.Phone) {
		fmt.Println("error is while input phone, format phone not correct")
		handleResponse(w, http.StatusBadRequest, user.Phone)
		return
	}

	id, err := c.Store.UserRepo.Insert(user)

	if err != nil {
		log.Fatalln("error is while inserting", err.Error())
		return
	}

	resp, err := c.Store.UserRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}

func (c Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	user, err := c.Store.UserRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, user)
}

func (c Controller) GetUserList(w http.ResponseWriter, r *http.Request) {
	users, err := c.Store.UserRepo.GetList()
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, users)
}

func (c Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	us := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&us); err != nil {
		fmt.Println("error is while decoding data", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if us.ID != id {
		fmt.Println("car ID not mismatch")
		handleResponse(w, http.StatusBadRequest, us.ID)
		return
	}

	if !check.PhoneNumber(us.Phone) {
		fmt.Println("error is while input phone, format phone not correct")
		handleResponse(w, http.StatusBadRequest, us.Phone)
		return
	}

	if err := c.Store.UserRepo.Update(us); err != nil {
		fmt.Println("error is while updating data", err.Error())
		return
	}

	resp, err := c.Store.UserRepo.GetByID(id)

	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	if err := c.Store.UserRepo.Delete(id); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return
	}

	handleResponse(w, http.StatusOK, id)
}
