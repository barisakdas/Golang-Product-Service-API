package entity

import (
	. "HB_Task/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Campaign struct {
	ID                     int
	ProductCode            string
	CampaignName           string
	Duration               float64
	PriceManupulationLimit float64
	AverageItemPrice       float64
	TargetSalesCount       int
	Turnover               int
	Status                 bool
}

func (c Campaign) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Campaign", "Migrate", err.Error())
	}

	db.AutoMigrate(&c)
}

func (c Campaign) Add() (Campaign, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Campaign", "Add", err.Error())
		return c, err
	}

	db.Create(&c)
	return c, nil
}

func (c Campaign) Get(where ...interface{}) (Campaign, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Campaign", "Get", err.Error())
		return c, err
	}

	db.First(&c, where...)
	return c, nil
}

func (c Campaign) GetAll(where ...interface{}) ([]Campaign, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Campaign", "GetAll", err.Error())
		return nil, err
	}

	var compaigns []Campaign
	db.Find(&compaigns, where...)
	return compaigns, nil
}

func (c Campaign) Update(data Campaign) (Campaign, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Campaign", "Update", err.Error())
		return c, err
	}

	db.Model(&c).Updates(data)
	return c, nil
}

func (c Campaign) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Campaign", "Update", err.Error())
		return err
	}

	db.Delete(&c)
	return nil
}
