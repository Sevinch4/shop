package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shop/models"
	"shop/storage/postgres"
)

type Controller struct {
	Store postgres.Store
}

func New(store postgres.Store) Controller {
	return Controller{Store: store}
}

func handleResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := models.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "succes"
	case code < 500:
		resp.Description = "bad request"
	default:
		resp.Description = "internal server error"
	}

	resp.StatusCode = statusCode
	resp.Data = data

	js, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error while marshalling resp", err.Error())
		return
	}

	w.WriteHeader(statusCode)
	w.Write(js)
}
