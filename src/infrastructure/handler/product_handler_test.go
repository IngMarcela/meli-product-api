package handler_test

import (
	"encoding/json"
	"errors"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"meli-product-api/internal/application/queries"
	"meli-product-api/internal/domain/model"
	"meli-product-api/src/infrastructure/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var testLogger *log.Logger

func init() {
	testLogger = log.New(os.Stdout, "[TEST] ", log.LstdFlags)
}

type mockProductService struct {
	product *model.Product
	err     error
}

func (m *mockProductService) GetByID(id string) (*model.Product, error) {
	return m.product, m.err
}

func TestGetProductByID_Success(t *testing.T) {
	// Arrange
	app := fiber.New()
	expectedProduct := &model.Product{ID: "123", Title: "Test Product"}
	mockService := &mockProductService{product: expectedProduct}
	query := queries.NewGetProductByIDQuery(mockService)
	productHandler := handler.NewProductHandler(query, testLogger)
	app.Get("/products/:id", productHandler.GetProductByID)

	// Act
	req := httptest.NewRequest("GET", "/products/123", nil)
	resp, _ := app.Test(req)

	// Assert
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var response model.Product
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct.ID, response.ID)
	assert.Equal(t, expectedProduct.Title, response.Title)
}

func TestGetProductByID_NotFound(t *testing.T) {
	// Arrange
	app := fiber.New()
	mockService := &mockProductService{err: errors.New("not found")}
	query := queries.NewGetProductByIDQuery(mockService)
	productHandler := handler.NewProductHandler(query, testLogger)
	app.Get("/products/:id", productHandler.GetProductByID)

	// Act
	req := httptest.NewRequest("GET", "/products/999", nil)
	resp, _ := app.Test(req)

	// Assert
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	var response map[string]string
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Producto no encontrado", response["error"])
}

func TestGetProductByID_EmptyID(t *testing.T) {
	// Arrange
	app := fiber.New()
	mockService := &mockProductService{}
	query := queries.NewGetProductByIDQuery(mockService)
	productHandler := handler.NewProductHandler(query, testLogger)
	app.Get("/products/:id", productHandler.GetProductByID)

	// Act
	req := httptest.NewRequest("GET", "/products/", nil)
	resp, _ := app.Test(req)

	// Assert
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}
