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

func GetAllCampaigns(w http.ResponseWriter, r *http.Request) {
	data, err := application.CampaignApplication{}.GetAllCampaign()
	if err != nil {
		log.LogJson("Error", "CampaignHandler", "GetAllCampaigns", err.Error())
	}
	if data == nil {
		util.HttpNotFound(w)
		return
	}
	jsonResp, _ := json.Marshal(data)
	w.Write(jsonResp)
}

func GetCurrentCampaign(w http.ResponseWriter, r *http.Request) {
	data, err := application.CampaignApplication{}.GetAllCampaign()
	if err != nil {
		log.LogJson("Error", "CampaignHandler", "GetCurrentCampaign", err.Error())
	}
	if data == nil {
		util.HttpNotFound(w)
		return
	}
	jsonResp, _ := json.Marshal(data)
	w.Write(jsonResp)
}

func GetCampaignsByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	data, err := application.CampaignApplication{}.GetCampaignByName(name)
	if err != nil {
		log.LogJson("Error", "CampaignHandler", "GetCampaignsByName", err.Error())
	}
	if (repository.CampaignDTO{} == data) {
		util.HttpNotFound(w)
		return
	}

	jsonResp, _ := json.Marshal(data)
	w.Write(jsonResp)
}

func IncreaseCampaignHour(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	hour := r.URL.Query().Get("hour")
	resp, err := application.CampaignApplication{}.IncreaseCampaignDuration(hour, name)
	if err != nil {
		log.LogJson("Error", "CampaignHandler", "IncreaseCampaignHour", err.Error())
	}
	if resp == "" {
		util.HttpNotFound(w)
		return
	}

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func AddNewCampaign(w http.ResponseWriter, r *http.Request) {
	var campaignDTO repository.CampaignDTO
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &campaignDTO)

	data, err := application.CampaignApplication{}.AddCampaign(&campaignDTO)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Data inserted: " + data))
}

func UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaignDTO repository.CampaignDTO
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &campaignDTO)

	application.CampaignApplication{}.UpdateCampaign(&campaignDTO)
	w.Write([]byte("Data updated."))
}

func DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	application.CampaignApplication{}.DeleteCampaign(name)
	w.Write([]byte("Deleting operation is success"))
}
