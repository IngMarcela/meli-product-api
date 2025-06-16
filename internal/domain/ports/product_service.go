package ports

import "meli-product-api/internal/domain/model"

// ProductService define los métodos que debe implementar un servicio de productos
type ProductService interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id string) (*model.Product, error)
}
