package main

import (
	"log"

	"meli-product-api/infrastructure/config"
)

func main() {
	config.InitLogger()
	config.Logger.Println("Iniciando servicio...")

	cfg := config.LoadConfig()

	app, err := config.SetupApplication(cfg)
	if err != nil {
		log.Fatalf("Error configurando la aplicación: %v", err)
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
