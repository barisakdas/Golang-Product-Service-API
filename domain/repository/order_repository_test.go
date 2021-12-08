package repository

import (
	"reflect"
	"testing"
)

func TestOrderRepository_GetAllOrder(t *testing.T) {
	tests := []struct {
		name    string
		o       OrderRepository
		want    []OrderDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.GetAllOrder()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepository.GetAllOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderRepository.GetAllOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderRepository_GetOrderByProductCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		o       OrderRepository
		args    args
		want    OrderDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.GetOrderByProductCode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepository.GetOrderByProductCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderRepository.GetOrderByProductCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderRepository_AddOrder(t *testing.T) {
	type args struct {
		orderDto *OrderDTO
	}
	tests := []struct {
		name    string
		o       OrderRepository
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.AddOrder(tt.args.orderDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepository.AddOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OrderRepository.AddOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderRepository_UpdateOrder(t *testing.T) {
	type args struct {
		orderDto *OrderDTO
	}
	tests := []struct {
		name    string
		o       OrderRepository
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.UpdateOrder(tt.args.orderDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepository.UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OrderRepository.UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderRepository_DeleteOrder(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		o       OrderRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.DeleteOrder(tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("OrderRepository.DeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
