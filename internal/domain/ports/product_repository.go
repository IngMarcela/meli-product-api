package ports

import "meli-product-api/internal/domain/model"

// ProductRepository define los métodos que debe implementar un repositorio de productos
type ProductRepository interface {
	GetByID(id string) (*model.Product, error)
}
