package repository

import (
	"reflect"
	"testing"
)

func TestCampaignRepository_GetAllCampaign(t *testing.T) {
	tests := []struct {
		name    string
		c       CampaignRepository
		want    []CampaignDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetAllCampaign()
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignRepository.GetAllCampaign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CampaignRepository.GetAllCampaign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignRepository_GetCampaignByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		c       CampaignRepository
		args    args
		want    CampaignDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetCampaignByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignRepository.GetCampaignByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CampaignRepository.GetCampaignByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignRepository_AddCampaign(t *testing.T) {
	type args struct {
		campaignDto *CampaignDTO
	}
	tests := []struct {
		name    string
		c       CampaignRepository
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.AddCampaign(tt.args.campaignDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignRepository.AddCampaign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CampaignRepository.AddCampaign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignRepository_UpdateCampaign(t *testing.T) {
	type args struct {
		campaignDto *CampaignDTO
	}
	tests := []struct {
		name    string
		c       CampaignRepository
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.UpdateCampaign(tt.args.campaignDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignRepository.UpdateCampaign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CampaignRepository.UpdateCampaign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignRepository_IncreaseCampaignDuration(t *testing.T) {
	type args struct {
		hour         string
		campaignName string
	}
	tests := []struct {
		name    string
		c       CampaignRepository
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.IncreaseCampaignDuration(tt.args.hour, tt.args.campaignName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignRepository.IncreaseCampaignDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CampaignRepository.IncreaseCampaignDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignRepository_DeleteCampaign(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		c       CampaignRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.DeleteCampaign(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("CampaignRepository.DeleteCampaign() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
