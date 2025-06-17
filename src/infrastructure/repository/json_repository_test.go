package repository

import (
	"encoding/json"
	"os"
	"testing"

	"meli-product-api/internal/domain/model"

	"github.com/stretchr/testify/assert"
)

func TestGetByID_Success(t *testing.T) {
	products := []model.Product{
		{
			ID:          "1",
			Title:       "Producto Uno",
			Description: "Descripción del producto",
			Price:       100.0,
			Stock:       10,
		},
	}

	data, _ := json.Marshal(products)
	tmpFile := "test_by_id.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo, err := NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Producto Uno", result.Title)
	assert.Equal(t, "Descripción del producto", result.Description)
	assert.Equal(t, 100.0, result.Price)
	assert.Equal(t, 10, result.Stock)
}

func TestGetByID_NotFound(t *testing.T) {
	products := []model.Product{
		{ID: "1", Title: "Producto Uno"},
	}

	data, _ := json.Marshal(products)
	tmpFile := "test_by_id_nf.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo, err := NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetByID("999")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no encontrado")
}

func TestGetByID_InvalidJSON(t *testing.T) {
	invalidJSON := `{"invalid": json}`
	tmpFile := "invalid.json"
	os.WriteFile(tmpFile, []byte(invalidJSON), 0644)
	defer os.Remove(tmpFile)

	_, err := NewJSONProductRepository(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error parseando JSON")
}

func TestGetByID_EmptyList(t *testing.T) {
	data := []model.Product{}
	jsonData, _ := json.Marshal(data)

	tmpFile := "empty_products.json"
	os.WriteFile(tmpFile, jsonData, 0644)
	defer os.Remove(tmpFile)

	repo, err := NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetByID("123")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no encontrado")
}

func TestGetByID_MultipleProducts(t *testing.T) {
	products := []model.Product{
		{ID: "1", Title: "Producto Uno"},
		{ID: "2", Title: "Producto Dos"},
		{ID: "3", Title: "Producto Tres"},
	}

	data, _ := json.Marshal(products)
	tmpFile := "multiple_products.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo, err := NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetByID("2")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Producto Dos", result.Title)
}

func TestGetAll_Success(t *testing.T) {
	products := []model.Product{
		{ID: "1", Title: "Producto Uno", Price: 100.0},
		{ID: "2", Title: "Producto Dos", Price: 200.0},
		{ID: "3", Title: "Producto Tres", Price: 300.0},
	}

	data, _ := json.Marshal(products)
	tmpFile := "test_get_all.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo, err := NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, result, 3)

	// Create a map for easier lookup
	resultMap := make(map[string]model.Product)
	for _, p := range result {
		resultMap[p.ID] = p
	}

	// Check each product by ID
	assert.Equal(t, "Producto Uno", resultMap["1"].Title)
	assert.Equal(t, 100.0, resultMap["1"].Price)
	assert.Equal(t, "Producto Dos", resultMap["2"].Title)
	assert.Equal(t, 200.0, resultMap["2"].Price)
	assert.Equal(t, "Producto Tres", resultMap["3"].Title)
	assert.Equal(t, 300.0, resultMap["3"].Price)
}

func TestGetAll_EmptyList(t *testing.T) {
	products := []model.Product{}
	data, _ := json.Marshal(products)
	tmpFile := "test_get_all_empty.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo, err := NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestGetAll_WithDuplicateIDs(t *testing.T) {
	products := []model.Product{
		{ID: "1", Title: "Producto Uno"},
		{ID: "1", Title: "Producto Uno Duplicado"},
		{ID: "2", Title: "Producto Dos"},
	}

	data, _ := json.Marshal(products)
	tmpFile := "test_get_all_duplicates.json"
	os.WriteFile(tmpFile, data, 0644)
	defer os.Remove(tmpFile)

	repo, err := NewJSONProductRepository(tmpFile)
	assert.NoError(t, err)
	result, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	for _, p := range result {
		if p.ID == "1" {
			assert.Equal(t, "Producto Uno Duplicado", p.Title)
		}
	}
}
