package application

import "HB_Task/domain/repository"

type OrderApplication struct{}

var orderRepository = repository.NewOrderRepository()

func (oa OrderApplication) GetAllOrder() ([]repository.OrderDTO, error) {
	return orderRepository.GetAllOrder()
}

func (oa OrderApplication) GetOrderByProductCode(code string) (repository.OrderDTO, error) {
	return orderRepository.GetOrderByProductCode(code)
}

func (oa OrderApplication) AddOrder(orderDto *repository.OrderDTO) (string, error) {
	return orderRepository.AddOrder(orderDto)
}

func (oa OrderApplication) UpdateOrder(orderDto *repository.OrderDTO) (string, error) {
	return orderRepository.UpdateOrder(orderDto)
}

func (oa OrderApplication) DeleteOrder(code string) error {
	return orderRepository.DeleteOrder(code)
}
