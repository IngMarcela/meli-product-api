package main

import (
	"log"

	"meli-product-api/infrastructure/config"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitLogger()
	config.Logger.Println("Iniciando servicio...")

	cfg := config.LoadConfig()

	// Configurar la aplicación
	app, err := config.SetupApplication(cfg)
	if err != nil {
		log.Fatalf("Error configurando la aplicación: %v", err)
	}

	// Middleware
	app.Use(cors.New())

	log.Printf("🚀 Servidor corriendo en http://localhost:%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
