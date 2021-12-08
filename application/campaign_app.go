package application

import "HB_Task/domain/repository"

type CampaignApplication struct{}

var campaignRepository = repository.NewCampaignRepository()

func (ca CampaignApplication) GetAllCampaign() ([]repository.CampaignDTO, error) {
	return campaignRepository.GetAllCampaign()
}

func (ca CampaignApplication) GetCampaignByName(name string) (repository.CampaignDTO, error) {
	return campaignRepository.GetCampaignByName(name)
}

func (ca CampaignApplication) AddCampaign(campaignDto *repository.CampaignDTO) (string, error) {
	return campaignRepository.AddCampaign(campaignDto)
}

func (ca CampaignApplication) UpdateCampaign(campaignDto *repository.CampaignDTO) (string, error) {
	return campaignRepository.UpdateCampaign(campaignDto)
}

func (ca CampaignApplication) IncreaseCampaignDuration(hour, campaignName string) (string, error) {
	return campaignRepository.IncreaseCampaignDuration(hour, campaignName)
}

func (ca CampaignApplication) DeleteCampaign(name string) error {
	return campaignRepository.DeleteCampaign(name)
}
