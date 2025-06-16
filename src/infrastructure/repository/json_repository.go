package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"meli-product-api/internal/domain/model"
	"meli-product-api/internal/domain/ports"
)

// jsonProductRepository implementa el puerto ProductRepository
type jsonProductRepository struct {
	filePath string
}

// NewJSONProductRepository crea una nueva instancia de jsonProductRepository
func NewJSONProductRepository(path string) ports.ProductRepository {
	return &jsonProductRepository{
		filePath: path,
	}
}

// GetAll obtiene todos los productos del archivo JSON
func (r *jsonProductRepository) GetAll() ([]model.Product, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo JSON: %w", err)
	}

	var products []model.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, fmt.Errorf("error parseando JSON: %w", err)
	}

	return products, nil
}

// GetByID busca un producto por ID
func (r *jsonProductRepository) GetByID(id string) (*model.Product, error) {
	products, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("producto con ID %s no encontrado", id)
}
