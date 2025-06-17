package handler

import (
	"log"
	"meli-product-api/internal/application/queries"

	"github.com/gofiber/fiber/v2"
)

// ProductHandler maneja las peticiones HTTP relacionadas con productos
type ProductHandler struct {
	productService *queries.GetProductByIDQuery
	logger         *log.Logger
}

// NewProductHandler crea una nueva instancia de ProductHandler
func NewProductHandler(productService *queries.GetProductByIDQuery, logger *log.Logger) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		logger:         logger,
	}
}

// GetProductByID maneja la petición GET /products/:id
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de producto requerido",
		})
	}

	h.logger.Printf("Recibida petición para producto con ID: %s", id)
	product, err := h.productService.Execute(id)
	if err != nil {
		h.logger.Printf("Error al obtener producto: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Producto no encontrado",
		})
	}

	return c.JSON(product)
}
