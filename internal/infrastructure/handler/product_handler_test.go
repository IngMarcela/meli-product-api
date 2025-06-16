package handler_test

import (
	"net/http/httptest"
	"testing"

	"meli-product-api/internal/application/queries"
	"meli-product-api/internal/application/usecases"
	"meli-product-api/internal/domain/model"
	"meli-product-api/internal/infrastructure/config"
       "meli-product-api/internal/infrastructure/handler"
       "meli-product-api/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.InitLogger()
}

type mockRepo struct{}

func (m *mockRepo) GetAll() ([]model.Product, error) {
	return []model.Product{
		{ID: "1", Title: "Galaxy S24", Description: "Smartphone Samsung"},
		{ID: "2", Title: "Poco X5", Description: "Smartphone Xiaomi"},
	}, nil
}

func (m *mockRepo) GetByID(id string) (*model.Product, error) {
	if id == "1" {
		return &model.Product{ID: "1", Title: "Galaxy S24", Description: "Smartphone Samsung"}, nil
	}
       return nil, repository.ErrProductNotFound
}

func TestGetProductByID_Success(t *testing.T) {
	app := fiber.New()
	mockQuery := queries.NewGetAllProductsQuery(&mockRepo{})
	mockID := queries.NewGetProductByIDQuery(&mockRepo{})
	search := usecases.NewSearchProductsUseCase(&mockRepo{})
	handler.RegisterProductRoutes(app, mockQuery, mockID, search)

	req := httptest.NewRequest("GET", "/products/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestGetProductByID_NotFound(t *testing.T) {
	app := fiber.New()
	mockQuery := queries.NewGetAllProductsQuery(&mockRepo{})
	mockID := queries.NewGetProductByIDQuery(&mockRepo{})
	search := usecases.NewSearchProductsUseCase(&mockRepo{})
	handler.RegisterProductRoutes(app, mockQuery, mockID, search)

	req := httptest.NewRequest("GET", "/products/999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestSearchProducts_Success(t *testing.T) {
	app := fiber.New()
	mockQuery := queries.NewGetAllProductsQuery(&mockRepo{})
	mockID := queries.NewGetProductByIDQuery(&mockRepo{})
	search := usecases.NewSearchProductsUseCase(&mockRepo{})
	handler.RegisterProductRoutes(app, mockQuery, mockID, search)

	req := httptest.NewRequest("GET", "/products/search?q=galaxy", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestSearchProducts_EmptyQuery(t *testing.T) {
	app := fiber.New()
	mockQuery := queries.NewGetAllProductsQuery(&mockRepo{})
	mockID := queries.NewGetProductByIDQuery(&mockRepo{})
	search := usecases.NewSearchProductsUseCase(&mockRepo{})
	handler.RegisterProductRoutes(app, mockQuery, mockID, search)

	req := httptest.NewRequest("GET", "/products/search", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestGetAllProducts_Success(t *testing.T) {
	app := fiber.New()
	mockQuery := queries.NewGetAllProductsQuery(&mockRepo{})
	mockID := queries.NewGetProductByIDQuery(&mockRepo{})
	search := usecases.NewSearchProductsUseCase(&mockRepo{})
	handler.RegisterProductRoutes(app, mockQuery, mockID, search)

	req := httptest.NewRequest("GET", "/products", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
