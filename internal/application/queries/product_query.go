package queries

import (
	"meli-product-api/internal/domain/model"
	"meli-product-api/internal/domain/ports"
)

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
