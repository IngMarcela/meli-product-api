package handler

import (
	"errors"
	"meli-product-api/internal/application/queries"
	"meli-product-api/internal/application/usecases"
	"meli-product-api/internal/infrastructure/config"
	"meli-product-api/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

// ProductHandler implementa el puerto ProductService
type ProductHandler struct {
	GetAllQuery   *queries.GetAllProductsQuery
	GetByIDQuery  *queries.GetProductByIDQuery
	SearchUseCase *usecases.SearchProductsUseCase
}

// NewProductHandler crea una nueva instancia de ProductHandler
func NewProductHandler(getAll *queries.GetAllProductsQuery, getByID *queries.GetProductByIDQuery, search *usecases.SearchProductsUseCase) *ProductHandler {
	return &ProductHandler{
		GetAllQuery:   getAll,
		GetByIDQuery:  getByID,
		SearchUseCase: search,
	}
}

// RegisterProductRoutes registra las rutas de los productos
func RegisterProductRoutes(app *fiber.App, getAll *queries.GetAllProductsQuery, getByID *queries.GetProductByIDQuery, search *usecases.SearchProductsUseCase) {
	handler := NewProductHandler(getAll, getByID, search)

	app.Get("/products", handler.GetAllProducts)
	app.Get("/products/search", handler.SearchProducts)
	app.Get("/products/:id", handler.GetProductByID)
}

// GetAllProducts implementa el puerto ProductRepository
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	config.Logger.Println("Obteniendo todos los productos")
	products, err := h.GetAllQuery.Execute()
	if err != nil {
		config.Logger.Printf("Error al obtener todos los productos: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	config.Logger.Printf("Productos obtenidos: %d", len(products))
	return c.JSON(products)
}

// GetProductByID implementa el puerto ProductRepository
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	config.Logger.Printf("Buscando producto con ID: %s", id)

	product, err := h.GetByIDQuery.Execute(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			config.Logger.Printf("Producto no encontrado: %v", err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Producto no encontrado"})
		}
		config.Logger.Printf("Error al buscar producto: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno"})
	}
	config.Logger.Printf("Producto encontrado: %v", product)
	return c.JSON(product)
}

// GET /products/search?q=xxx
func (h *ProductHandler) SearchProducts(c *fiber.Ctx) error {
	config.Logger.Println("Buscando productos")
	q := c.Query("q")
	if q == "" {
		config.Logger.Println("Falta el parámetro de búsqueda")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing query param 'q'"})
	}

	result, err := h.SearchUseCase.Execute(q)
	if err != nil {
		config.Logger.Printf("Error al buscar productos: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al buscar productos"})
	}

	config.Logger.Printf("Productos encontrados: %d", len(result))
	return c.JSON(result)
}
