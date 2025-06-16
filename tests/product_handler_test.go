package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"meli-product-api/internal/application/queries"
	"meli-product-api/internal/application/usecases"
	"meli-product-api/internal/domain/model"
	"meli-product-api/internal/infrastructure/config"
	"meli-product-api/internal/infrastructure/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.InitLogger()
}

type testRepo struct{}

func (r *testRepo) GetAll() ([]model.Product, error) {
	return []model.Product{
		{ID: "100", Title: "Producto Test"},
	}, nil
}

func (r *testRepo) GetByID(id string) (*model.Product, error) {
	if id == "100" {
		return &model.Product{ID: "100", Title: "Producto Test"}, nil
	}
	return nil, errors.New("product not found")
}

func TestGetAllProductsRoute(t *testing.T) {
	app := fiber.New()

	repo := &testRepo{}
	getAll := queries.NewGetAllProductsQuery(repo)
	getByID := queries.NewGetProductByIDQuery(repo)
	search := usecases.NewSearchProductsUseCase(repo)
	handler.RegisterProductRoutes(app, getAll, getByID, search)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error al hacer la request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado 200, recibido %d", resp.StatusCode)
	}

	var products []model.Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		t.Fatalf("error al parsear JSON: %v", err)
	}

	if len(products) != 1 || products[0].ID != "100" {
		t.Errorf("producto inesperado: %+v", products)
	}
}

func TestGetProductByID_Success(t *testing.T) {
	app := fiber.New()

	repo := &testRepo{}
	getAll := queries.NewGetAllProductsQuery(repo)
	getByID := queries.NewGetProductByIDQuery(repo)
	search := usecases.NewSearchProductsUseCase(repo)
	handler.RegisterProductRoutes(app, getAll, getByID, search)

	req := httptest.NewRequest(http.MethodGet, "/products/100", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error en request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado 200, recibido %d", resp.StatusCode)
	}

	var product model.Product
	json.NewDecoder(resp.Body).Decode(&product)

	if product.ID != "100" {
		t.Errorf("esperado producto con ID 100, recibido %s", product.ID)
	}
}

func TestGetProductByID_NotFound(t *testing.T) {
	app := fiber.New()

	repo := &testRepo{}
	getAll := queries.NewGetAllProductsQuery(repo)
	getByID := queries.NewGetProductByIDQuery(repo)
	search := usecases.NewSearchProductsUseCase(repo)
	handler.RegisterProductRoutes(app, getAll, getByID, search)

	req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error en request: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("esperado 404, recibido %d", resp.StatusCode)
	}
}

func TestSearchProducts_Success(t *testing.T) {
	app := fiber.New()

	repo := &testRepo{}
	mockQuery := queries.NewGetAllProductsQuery(repo)
	mockID := queries.NewGetProductByIDQuery(repo)
	search := usecases.NewSearchProductsUseCase(repo)

	handler.RegisterProductRoutes(app, mockQuery, mockID, search)

	req := httptest.NewRequest("GET", "/products/search?q=galaxy", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestSearchProducts_EmptyQuery(t *testing.T) {
	app := fiber.New()

	repo := &testRepo{}
	mockQuery := queries.NewGetAllProductsQuery(repo)
	mockID := queries.NewGetProductByIDQuery(repo)
	search := usecases.NewSearchProductsUseCase(repo)

	handler.RegisterProductRoutes(app, mockQuery, mockID, search)

	req := httptest.NewRequest("GET", "/products/search", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
