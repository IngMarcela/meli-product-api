package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"meli-product-api/internal/domain/model"
	"meli-product-api/internal/domain/ports"
)

type jsonProductRepository struct {
	products map[string]model.Product
}

// NewJSONProductRepository carga el JSON una vez al iniciar y construye un mapa para acceso eficiente por ID
func NewJSONProductRepository(path string) (ports.ProductRepository, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo JSON: %w", err)
	}

	var productList []model.Product
	if err := json.Unmarshal(data, &productList); err != nil {
		return nil, fmt.Errorf("error parseando JSON: %w", err)
	}

	productMap := make(map[string]model.Product)
	for _, p := range productList {
		productMap[p.ID] = p
	}

	return &jsonProductRepository{
		products: productMap,
	}, nil
}

// GetAll retorna todos los productos desde el mapa en memoria
func (r *jsonProductRepository) GetAll() ([]model.Product, error) {
	productList := make([]model.Product, 0, len(r.products))
	for _, p := range r.products {
		productList = append(productList, p)
	}
	return productList, nil
}

// GetByID busca el producto
func (r *jsonProductRepository) GetByID(id string) (*model.Product, error) {
	if p, exists := r.products[id]; exists {
		return &p, nil
	}
	return nil, fmt.Errorf("producto con ID %s no encontrado", id)
}
