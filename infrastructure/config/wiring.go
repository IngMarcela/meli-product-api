package config

import (
	"meli-product-api/internal/application/queries"
	"meli-product-api/src/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupApplication configura todas las dependencias de la aplicación y las rutas
func SetupApplication(cfg *Config) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		AppName: "Meli Product API",
	})

	app.Use(cors.New())

	productRepo, err := repository.NewJSONProductRepository(cfg.ProductFile)
	if err != nil {
		return nil, err
	}

	getByID := queries.NewGetProductByIDQuery(productRepo)

	RegisterRoutes(app, getByID, Logger)

	return app, nil
}
