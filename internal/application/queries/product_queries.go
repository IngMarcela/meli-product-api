package queries

import (
	"meli-product-service/internal/domain/model"
	"meli-product-service/internal/domain/ports"
)

// GetAllProductsQuery implementa el puerto ProductRepository
type GetAllProductsQuery struct {
	repo ports.ProductRepository
}

// NewGetAllProductsQuery crea una nueva instancia de GetAllProductsQuery
func NewGetAllProductsQuery(repo ports.ProductRepository) *GetAllProductsQuery {
	return &GetAllProductsQuery{repo: repo}
}

func (q *GetAllProductsQuery) Execute() ([]model.Product, error) {
	return q.repo.GetAll()
}

// GetProductByIDQuery implementa el puerto ProductRepository
type GetProductByIDQuery struct {
	repo ports.ProductRepository
}

// NewGetProductByIDQuery crea una nueva instancia de GetProductByIDQuery
func NewGetProductByIDQuery(repo ports.ProductRepository) *GetProductByIDQuery {
	return &GetProductByIDQuery{repo: repo}
}

func (q *GetProductByIDQuery) Execute(id string) (*model.Product, error) {
	return q.repo.GetByID(id)
}
