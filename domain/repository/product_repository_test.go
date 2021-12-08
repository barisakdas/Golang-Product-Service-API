package repository

import (
	"reflect"
	"testing"
)

func TestProductRepository_GetAllProducts(t *testing.T) {
	tests := []struct {
		name    string
		p       ProductRepository
		want    []ProductDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetAllProducts()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.GetAllProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepository.GetAllProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepository_GetProductByProductCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		p       ProductRepository
		args    args
		want    ProductDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetProductByProductCode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.GetProductByProductCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepository.GetProductByProductCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepository_AddProduct(t *testing.T) {
	type args struct {
		productDto *ProductDTO
	}
	tests := []struct {
		name    string
		p       ProductRepository
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.AddProduct(tt.args.productDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProductRepository.AddProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepository_UpdateProduct(t *testing.T) {
	type args struct {
		productDto *ProductDTO
	}
	tests := []struct {
		name    string
		p       ProductRepository
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.UpdateProduct(tt.args.productDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProductRepository.UpdateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepository_DeleteProduct(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		p       ProductRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.DeleteProduct(tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
