package config

import (
	"log"
	"meli-product-api/internal/application/queries"
	"meli-product-api/src/infrastructure/handler"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes configura todas las rutas de la aplicación
func RegisterRoutes(app *fiber.App, getByID *queries.GetProductByIDQuery, logger *log.Logger) {
	productHandler := handler.NewProductHandler(getByID, logger)

	// Rutas de productos
	app.Get("/products/:id", productHandler.GetProductByID)
}
