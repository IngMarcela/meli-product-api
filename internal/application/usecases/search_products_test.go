package usecases_test

import (
	"testing"

	"meli-product-api/internal/application/usecases"
	"meli-product-api/internal/domain/model"

	"github.com/stretchr/testify/assert"
)

type mockRepo struct{}

func (m *mockRepo) GetAll() ([]model.Product, error) {
	return []model.Product{
		{ID: "1", Title: "Galaxy S24", Description: "Teléfono Samsung"},
		{ID: "2", Title: "Xiaomi Poco", Description: "Potente y económico"},
	}, nil
}

func (m *mockRepo) GetByID(id string) (*model.Product, error) {
	return nil, nil
}

func TestSearchProducts_Execute_Match(t *testing.T) {
	uc := usecases.NewSearchProductsUseCase(&mockRepo{})

	result, err := uc.Execute("galaxy")

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].ID)
}
