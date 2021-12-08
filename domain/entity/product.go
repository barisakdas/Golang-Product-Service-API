package entity

import (
	. "HB_Task/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID                       int
	ProductCode, ProductName string
	Price                    float64
	Stock                    int
}

func (p Product) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Product", "Migrate", err.Error())
	}

	db.AutoMigrate(p)
}

func (p Product) Add() (Product, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Product", "Add", err.Error())
	}

	db.Create(p)
	return p, nil
}

func (p Product) Get(where ...interface{}) (Product, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Product", "Get", err.Error())
		return p, err
	}

	db.First(p, where...)
	return p, nil
}
func (p Product) GetAll(where ...interface{}) ([]Product, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Product", "GetAll", err.Error())
		return nil, err
	}

	var products []Product
	db.Find(products, where...)
	return products, err
}

func (p Product) Update(data Product) (Product, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Product", "Update", err.Error())
		return p, err
	}

	db.Model(p).Updates(data)
	return p, nil
}

func (p Product) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Product", "Update", err.Error())
		return err
	}

	db.Delete(p)
	return nil
}
