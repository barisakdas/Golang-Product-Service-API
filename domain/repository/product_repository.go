package repository

import (
	"HB_Task/domain/entity"
	"HB_Task/log"
)

type IProductRepository interface {
	GetAllProducts() ([]ProductDTO, error)
	GetProductByProductCode(code string) (ProductDTO, error)
	AddProduct(productDto *ProductDTO) (string, error)
	UpdateProduct(productDto *ProductDTO) (string, error)
	DeleteProduct(code string) error
}

type ProductRepository struct{}

type ProductDTO struct {
	ProductCode, ProductName string
	Price                    float64
	Stock                    int
}

func NewProductRepository() IProductRepository {
	return &ProductRepository{}
}

func (p ProductRepository) GetAllProducts() ([]ProductDTO, error) {
	data, err := entity.Product{}.GetAll()
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "GetAllProduct", err.Error())
		return nil, err
	}
	var dataDtos []ProductDTO

	for _, value := range data {
		newData := ProductDTO{
			ProductCode: value.ProductCode,
			ProductName: value.ProductName,
			Price:       value.Price,
			Stock:       value.Stock,
		}
		dataDtos = append(dataDtos, newData)
	}
	return dataDtos, nil
}

func (p ProductRepository) GetProductByProductCode(code string) (ProductDTO, error) {
	data, err := entity.Product{}.Get("product_code = ?", code)
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "GetProductByProductCode", err.Error())
		return ProductDTO{}, err
	}

	return ProductDTO{
		ProductCode: data.ProductCode,
		ProductName: data.ProductName,
		Price:       data.Price,
		Stock:       data.Stock,
	}, nil
}

func (p ProductRepository) AddProduct(productDto *ProductDTO) (string, error) {
	result, err := entity.Product{
		ProductCode: productDto.ProductCode,
		ProductName: productDto.ProductName,
		Price:       productDto.Price,
		Stock:       productDto.Stock,
	}.Add()
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "AddProduct", err.Error())
		return "", err
	}
	return result.ProductCode, nil
}

func (p ProductRepository) UpdateProduct(productDto *ProductDTO) (string, error) {
	data, err := entity.Product{}.Get("product_code = ?", productDto.ProductCode)
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "UpdateProduct-Get", err.Error())
		return "", err
	}

	result, err := data.Update(entity.Product{
		ProductCode: productDto.ProductCode,
		ProductName: productDto.ProductName,
		Price:       productDto.Price,
		Stock:       productDto.Stock,
	})

	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "UpdateProduct", err.Error())
		return "", err
	}

	return result.ProductCode, nil
}

func (p ProductRepository) DeleteProduct(code string) error {
	data, err := entity.Product{}.Get("product_code = ?", code)
	if err != nil {
		log.LogJson("Error", "Repository/ProductRepository", "DeleteProduct", err.Error())
		return err
	}

	return data.Delete()
}
