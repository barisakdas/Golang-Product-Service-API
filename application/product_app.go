package application

import "HB_Task/domain/repository"

type ProductApplication struct{}

var productRepository = repository.NewProductRepository()

func (pa ProductApplication) GetAllProducts() ([]repository.ProductDTO, error) {
	return productRepository.GetAllProducts()
}

func (pa ProductApplication) GetProductByProductCode(code string) (repository.ProductDTO, error) {
	return productRepository.GetProductByProductCode(code)
}

func (pa ProductApplication) AddProduct(productDto *repository.ProductDTO) (string, error) {
	return productRepository.AddProduct(productDto)
}

func (pa ProductApplication) UpdateProduct(productDto *repository.ProductDTO) (string, error) {
	return productRepository.UpdateProduct(productDto)
}

func (pa ProductApplication) DeleteProduct(code string) error {
	return productRepository.DeleteProduct(code)
}
