package config

import (
	"meli-product-api/internal/application/queries"
	"meli-product-api/src/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

// SetupApplication configura todas las dependencias de la aplicación
func SetupApplication(cfg *Config) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		AppName: "Meli Product API",
	})

	productRepo, err := repository.NewJSONProductRepository(cfg.ProductFile)
	if err != nil {
		return nil, err
	}

	getByID := queries.NewGetProductByIDQuery(productRepo)

	RegisterRoutes(app, getByID, Logger)

	return app, nil
}
