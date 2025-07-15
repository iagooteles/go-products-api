package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecaseInterface interface {
	GetProducts() ([]model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	GetProductById(productId int) (*model.Product, error)
	UpdateProduct(product model.Product) (model.Product, error)
	DeleteProduct(productId int) (bool, error)
}

type ProductUsecase struct {
	repository repository.ProductRepository
}

var _ ProductUsecaseInterface = (*ProductUsecase)(nil)


func NewProductUsecase(repo repository.ProductRepository) ProductUsecaseInterface {
	return &ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) GetProductById(productId int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(productId)

	if err != nil {
		return nil, err
	}
	
	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.UpdateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) DeleteProduct(productId int) (bool, error) {
	deleted, err := pu.repository.DeleteProduct(productId)

	if err != nil {
		return false, err
	}
	
	return deleted, nil
}