package config

import (
	"os"
)

// Config define la configuración de la aplicación
type Config struct {
	Port        string
	ProductFile string
}

// LoadConfig carga la configuración desde variables de entorno o valores por defecto
func LoadConfig() *Config {
	port := getEnv("APP_PORT", "3000")
	file := getEnv("PRODUCT_FILE", "data/products.json")

	return &Config{
		Port:        port,
		ProductFile: file,
	}
}

// getEnv obtiene el valor de una variable de entorno o devuelve un valor por defecto
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
