package config

import (
	"log"
	"os"
)

// Logger es el logger global de la aplicación
var Logger *log.Logger

// InitLogger inicializa el logger global
func InitLogger() {
	Logger = log.New(os.Stdout, "[PRODUCT-SERVICE] ", log.LstdFlags)
}

// Config define la configuración de la aplicación
type Config struct {
	Port        string
	ProductFile string
}

// LoadConfig carga la configuración desde variables de entorno o valores por defecto
func LoadConfig() *Config {
	return &Config{
		Port:        "3000",
		ProductFile: "data/products.json",
	}
}
