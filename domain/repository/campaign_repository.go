package repository

import (
	"HB_Task/domain/entity"
	"HB_Task/log"
	"errors"
	"strconv"
)

type ICampaignRepository interface {
	GetAllCampaign() ([]CampaignDTO, error)
	GetCampaignByName(name string) (CampaignDTO, error)
	AddCampaign(campaignDto *CampaignDTO) (string, error)
	UpdateCampaign(campaignDto *CampaignDTO) (string, error)
	IncreaseCampaignDuration(hour, campaignName string) (string, error)
	DeleteCampaign(name string) error
}

type CampaignRepository struct{}

func NewCampaignRepository() ICampaignRepository {
	return &CampaignRepository{}
}

type CampaignDTO struct {
	ProductCode            string
	CampaignName           string
	Duration               float64
	PriceManupulationLimit float64
	AverageItemPrice       float64
	TargetSalesCount       int
	Turnover               int
	Status                 bool
}

func (c CampaignRepository) GetAllCampaign() ([]CampaignDTO, error) {
	data, err := entity.Campaign{}.GetAll()
	if err != nil {
		log.LogJson("Error", "Repository/CampaignRepository", "GetAllCampaign", err.Error())
		return nil, err
	}

	var dataDtos []CampaignDTO

	for _, value := range data {
		newData := CampaignDTO{
			ProductCode:            value.ProductCode,
			CampaignName:           value.CampaignName,
			Duration:               value.Duration,
			PriceManupulationLimit: value.PriceManupulationLimit,
			AverageItemPrice:       value.AverageItemPrice,
			TargetSalesCount:       value.TargetSalesCount,
			Turnover:               value.Turnover,
			Status:                 value.Status,
		}
		dataDtos = append(dataDtos, newData)
	}
	return dataDtos, err
}

func (c CampaignRepository) GetCampaignByName(name string) (CampaignDTO, error) {
	data, err := entity.Campaign{}.Get("campaign_name = ?", name)
	if err != nil {
		log.LogJson("Error", "Repository/CampaignRepository", "GetCampaignByName", err.Error())
		return CampaignDTO{}, err
	}
	return CampaignDTO{
		ProductCode:            data.ProductCode,
		CampaignName:           data.CampaignName,
		Duration:               data.Duration,
		PriceManupulationLimit: data.PriceManupulationLimit,
		AverageItemPrice:       data.AverageItemPrice,
		TargetSalesCount:       data.TargetSalesCount,
		Turnover:               data.Turnover,
		Status:                 data.Status,
	}, nil
}

func (c CampaignRepository) AddCampaign(campaignDto *CampaignDTO) (string, error) {
	data, _ := entity.Campaign{}.Get("campaign_name = ? and product_code = ?", campaignDto.CampaignName, campaignDto.ProductCode)
	if (entity.Campaign{}) == data {
		result, err := entity.Campaign{
			ProductCode:            campaignDto.ProductCode,
			CampaignName:           campaignDto.CampaignName,
			Duration:               campaignDto.Duration,
			PriceManupulationLimit: campaignDto.PriceManupulationLimit,
			AverageItemPrice:       campaignDto.AverageItemPrice,
			TargetSalesCount:       campaignDto.TargetSalesCount,
			Turnover:               campaignDto.Turnover,
			Status:                 campaignDto.Status,
		}.Add()
		if err != nil {
			log.LogJson("Error", "Repository/CampaignRepository", "AddCampaign", err.Error())
			return "", err
		}
		return result.CampaignName, nil
	}
	return "", errors.New("Data already exist")
}

func (c CampaignRepository) UpdateCampaign(campaignDto *CampaignDTO) (string, error) {
	data, err := entity.Campaign{}.Get("campaign_name = ?", campaignDto.CampaignName)
	if err != nil {
		log.LogJson("Error", "Repository/CampaignRepository", "GetCampaignByName-Get", err.Error())
		return "", err
	}

	result, err := data.Update(entity.Campaign{
		ProductCode:            campaignDto.ProductCode,
		CampaignName:           campaignDto.CampaignName,
		Duration:               campaignDto.Duration,
		PriceManupulationLimit: campaignDto.PriceManupulationLimit,
		AverageItemPrice:       campaignDto.AverageItemPrice,
		TargetSalesCount:       campaignDto.TargetSalesCount,
		Turnover:               campaignDto.Turnover,
		Status:                 campaignDto.Status,
	})

	if err != nil {
		log.LogJson("Error", "Repository/CampaignRepository", "UpdateCampaign", err.Error())
		return "", err
	}
	return result.CampaignName, nil
}

func (c CampaignRepository) IncreaseCampaignDuration(hour, campaignName string) (string, error) {
	data, err := entity.Campaign{}.Get("campaign_name = ?", campaignName)
	additionalTime, _ := strconv.ParseFloat(hour, 32)
	if err != nil {
		log.LogJson("Error", "Repository/CampaignRepository", "GetCampaignByName-Get", err.Error())
		return "", err
	}

	result, err := data.Update(entity.Campaign{
		Duration: data.Duration + additionalTime,
	})

	if err != nil {
		log.LogJson("Error", "Repository/CampaignRepository", "UpdateCampaign", err.Error())
		return "", err
	}
	return result.CampaignName, nil
}

func (c CampaignRepository) DeleteCampaign(name string) error {
	data, err := entity.Campaign{}.Get("campaign_name = ?", name)
	if err != nil {
		log.LogJson("Error", "Repository/CampaignRepository", "GetCampaignByName-Get", err.Error())
		return err
	}
	return data.Delete()
}
