package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"meli-product-service/internal/infrastructure/config"
	"meli-product-service/internal/infrastructure/handler"
	"meli-product-service/internal/infrastructure/repository"

	"meli-product-service/internal/application/queries"
	"meli-product-service/internal/application/usecases"
)

func main() {
	config.InitLogger()
	config.Logger.Println("Iniciando servicio...")

	cfg := config.LoadConfig()

	app := fiber.New()
	app.Use(cors.New())

	productRepo := repository.NewJSONProductRepository(cfg.ProductFile)

	getAll := queries.NewGetAllProductsQuery(productRepo)
	getByID := queries.NewGetProductByIDQuery(productRepo)
	search := usecases.NewSearchProductsUseCase(productRepo)

	handler.RegisterProductRoutes(app, getAll, getByID, search)

	log.Printf("🚀 Servidor corriendo en http://localhost:%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
