package repository_test

import (
	"encoding/json"
	"meli-product-api/internal/domain/model"
	"meli-product-api/src/infrastructure/repository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByID_Found(t *testing.T) {
	products := []model.Product{
		{ID: "1", Title: "Producto Uno"},
	}

	data, _ := json.Marshal(products)
	tmpFile := "test_by_id.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo, err := repository.NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Producto Uno", result.Title)
}

func TestGetByID_NotFound(t *testing.T) {
	products := []model.Product{
		{ID: "1", Title: "Producto Uno"},
	}

	data, _ := json.Marshal(products)
	tmpFile := "test_by_id_nf.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo, err := repository.NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetByID("999")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no encontrado")
}

func TestGetByID_EmptyList(t *testing.T) {
	data := []model.Product{} // lista vacía
	jsonData, _ := json.Marshal(data)

	tmpFile := "empty_products.json"
	os.WriteFile(tmpFile, jsonData, 0644)
	defer os.Remove(tmpFile)

	repo, err := repository.NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetByID("123")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no encontrado")
}
