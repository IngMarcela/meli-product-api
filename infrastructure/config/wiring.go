package config

import (
	"meli-product-api/internal/application/queries"
	"meli-product-api/src/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

// SetupApplication configura todas las dependencias de la aplicación
func SetupApplication(cfg *Config) (*fiber.App, error) {
	// Crear instancia de Fiber
	app := fiber.New(fiber.Config{
		AppName: "Meli Product API",
	})

	// Inicializar repositorios
	productRepo, err := repository.NewJSONProductRepository(cfg.ProductFile)
	if err != nil {
		return nil, err
	}

	// Inicializar queries
	getByID := queries.NewGetProductByIDQuery(productRepo)

	// Configurar rutas
	RegisterRoutes(app, getByID, Logger)

	return app, nil
}
