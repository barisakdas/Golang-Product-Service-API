package entity

import (
	. "HB_Task/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID          int
	ProductCode string
	Quantity    int
	TotalPrice  float64
}

func (o Order) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Order", "Migrate", err.Error())
	}

	db.AutoMigrate(o)
}

func (o Order) Add() (Order, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Order", "Add", err.Error())
		return o, err
	}

	db.Create(o)
	return o, nil
}

func (o Order) Get(where ...interface{}) (Order, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Order", "Get", err.Error())
		return o, err
	}

	db.First(o, where...)
	return o, nil
}
func (o Order) GetAll(where ...interface{}) ([]Order, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Order", "GetAll", err.Error())
		return nil, err
	}

	var orders []Order
	db.Find(&orders, where...)
	return orders, nil
}

func (o Order) Update(data Order) (Order, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Order", "Update", err.Error())
		return o, err
	}

	db.Model(o).Updates(data)
	return o, nil
}

func (o Order) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		LogJson("Error", "Entity/Order", "Update", err.Error())
		return err
	}

	db.Delete(o)
	return nil
}
