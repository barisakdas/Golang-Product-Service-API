package presentation

import (
	"HB_Task/application"
	"HB_Task/domain/repository"
	"HB_Task/infrastructure/util"
	"HB_Task/log"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	data, err := application.OrderApplication{}.GetAllOrder()
	if err != nil {
		log.LogJson("Error", "OrderHandler", "GetAllOrder", err.Error())
	}
	if data == nil {
		util.HttpNotFound(w)
		return
	}
	jsonResp, _ := json.Marshal(data)
	w.Write(jsonResp)
}

func GetOrderByProductCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	data, err := application.OrderApplication{}.GetOrderByProductCode(code)
	if err != nil {
		log.LogJson("Error", "OrderHandler", "GetOrderByProductCode", err.Error())
	}
	if (repository.OrderDTO{} == data) {
		util.HttpNotFound(w)
		return
	}

	jsonResp, _ := json.Marshal(data)
	w.Write(jsonResp)
}

func AddNewOrder(w http.ResponseWriter, r *http.Request) {
	var orderDto repository.OrderDTO
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &orderDto)

	application.OrderApplication{}.AddOrder(&orderDto)
	w.Write([]byte("Data inserted."))
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDto repository.OrderDTO
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &orderDto)

	application.OrderApplication{}.UpdateOrder(&orderDto)
	w.Write([]byte("Data updated."))
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	application.OrderApplication{}.DeleteOrder(code)
	w.Write([]byte("Deleting operation is success"))
}
