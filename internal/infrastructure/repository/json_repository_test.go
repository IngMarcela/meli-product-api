package repository_test

import (
	"encoding/json"
	"meli-product-api/internal/domain/model"
	"meli-product-api/internal/infrastructure/repository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll_ReadsFileSuccessfully(t *testing.T) {
	// Crear archivo temporal
	data := []model.Product{
		{ID: "1", Title: "Producto 1"},
		{ID: "2", Title: "Producto 2"},
	}
	jsonData, _ := json.Marshal(data)

	tmpFile := "test_products.json"
	os.WriteFile(tmpFile, jsonData, 0644)
	defer os.Remove(tmpFile)

	repo := repository.NewJSONProductRepository(tmpFile)
	products, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Producto 1", products[0].Title)
}

func TestGetByID_Found(t *testing.T) {
	products := []model.Product{
		{ID: "1", Title: "Producto Uno"},
	}

	data, _ := json.Marshal(products)
	tmpFile := "test_by_id.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo := repository.NewJSONProductRepository(tmpFile)
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

	repo := repository.NewJSONProductRepository(tmpFile)
	result, err := repo.GetByID("999")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no encontrado")
}

func TestGetByID_FailsIfGetAllFails(t *testing.T) {
	repo := repository.NewJSONProductRepository("no_existe.json")

	result, err := repo.GetByID("1")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error leyendo archivo JSON")
}

func TestGetByID_EmptyList(t *testing.T) {
	data := []model.Product{} // lista vacía
	jsonData, _ := json.Marshal(data)

	tmpFile := "empty_products.json"
	os.WriteFile(tmpFile, jsonData, 0644)
	defer os.Remove(tmpFile)

	repo := repository.NewJSONProductRepository(tmpFile)
	result, err := repo.GetByID("123")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no encontrado")
}
