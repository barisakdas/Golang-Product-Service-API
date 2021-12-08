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

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	data, err := application.ProductApplication{}.GetAllProducts()
	if err != nil {
		log.LogJson("Error", "ProductHandler", "GetAllProducts", err.Error())
	}
	if data == nil {
		util.HttpNotFound(w)
		return
	}
	jsonResp, _ := json.Marshal(data)
	w.Write(jsonResp)
}

func GetProductByProductCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	data, err := application.ProductApplication{}.GetProductByProductCode(code)
	if err != nil {
		log.LogJson("Error", "CampaignHandler", "GetCampaignsByName", err.Error())
	}
	if (repository.ProductDTO{} == data) {
		util.HttpNotFound(w)
		return
	}

	jsonResp, _ := json.Marshal(data)
	w.Write(jsonResp)
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO repository.ProductDTO
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &productDTO)

	data, err := application.ProductApplication{}.AddProduct(&productDTO)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Data inserted: " + data))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO repository.ProductDTO
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &productDTO)

	data, err := application.ProductApplication{}.UpdateProduct(&productDTO)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Data updated: " + data))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	application.ProductApplication{}.DeleteProduct(code)
	w.Write([]byte("Deleting operation is success"))
}
