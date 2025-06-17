package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"meli-product-api/infrastructure/config"
	"meli-product-api/internal/application/queries"
	"meli-product-api/internal/domain/model"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.InitLogger()
}

type testRepo struct{}

func (r *testRepo) GetByID(id string) (*model.Product, error) {
	if id == "100" {
		return &model.Product{ID: "100", Title: "Producto Test"}, nil
	}
	return nil, errors.New("product not found")
}

func TestGetProductByID_Success(t *testing.T) {
	app := fiber.New()

	repo := &testRepo{}
	getByID := queries.NewGetProductByIDQuery(repo)
	config.RegisterRoutes(app, getByID, config.Logger)

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
	getByID := queries.NewGetProductByIDQuery(repo)
	config.RegisterRoutes(app, getByID, config.Logger)

	req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error en request: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("esperado 404, recibido %d", resp.StatusCode)
	}
}
