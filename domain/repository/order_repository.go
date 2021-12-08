package repository

import (
	"HB_Task/domain/entity"
	"HB_Task/log"
	"errors"
)

type IOrderRepository interface {
	GetAllOrder() ([]OrderDTO, error)
	GetOrderByProductCode(code string) (OrderDTO, error)
	AddOrder(orderDto *OrderDTO) (string, error)
	UpdateOrder(orderDto *OrderDTO) (string, error)
	DeleteOrder(code string) error
}

type OrderRepository struct{}

func NewOrderRepository() IOrderRepository {
	return &OrderRepository{}
}

type OrderDTO struct {
	ProductCode string
	Quantity    int
	TotalPrice  float64
}

func (o OrderRepository) GetAllOrder() ([]OrderDTO, error) {
	data, err := entity.Order{}.GetAll()
	if err != nil {
		log.LogJson("Error", "Repository/OrderRepository", "GetAllOrder", err.Error())
		return nil, err
	}
	var dataDtos []OrderDTO
	for _, value := range data {
		newData := OrderDTO{
			ProductCode: value.ProductCode,
			Quantity:    value.Quantity,
			TotalPrice:  value.TotalPrice,
		}
		dataDtos = append(dataDtos, newData)
	}
	return dataDtos, nil
}

func (o OrderRepository) GetOrderByProductCode(code string) (OrderDTO, error) {
	data, err := entity.Order{}.Get("product_code = ?", code)
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "GetOrderByProductCode", err.Error())
		return OrderDTO{}, err
	}
	return OrderDTO{
		ProductCode: data.ProductCode,
		Quantity:    data.Quantity,
		TotalPrice:  data.TotalPrice,
	}, nil
}

func (o OrderRepository) AddOrder(orderDto *OrderDTO) (string, error) {
	data, _ := entity.Order{}.Get("product_code = ?", orderDto.ProductCode)
	if (entity.Order{}) == data {
		result, err := entity.Order{
			ProductCode: orderDto.ProductCode,
			Quantity:    orderDto.Quantity,
			TotalPrice:  orderDto.TotalPrice,
		}.Add()
		if err != nil {
			log.LogJson("Error", "Repository/ProductRepository", "AddOrder", err.Error())
			return "", err
		}
		return result.ProductCode, nil
	}
	return "", errors.New("Data already exist")
}

func (o OrderRepository) UpdateOrder(orderDto *OrderDTO) (string, error) {
	data, err := entity.Order{}.Get("product_code = ?", orderDto.ProductCode)
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "UpdateOrder-Get", err.Error())
		return "", err
	}
	result, err := data.Update(entity.Order{
		ProductCode: orderDto.ProductCode,
		Quantity:    orderDto.Quantity,
		TotalPrice:  orderDto.TotalPrice,
	})
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "UpdateOrder", err.Error())
		return "", err
	}
	return result.ProductCode, nil
}

func (o OrderRepository) DeleteOrder(code string) error {
	data, err := entity.Order{}.Get("product_code = ?", code)
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "UpdateOrder-Get", err.Error())
		return err
	}
	return data.Delete()
}
