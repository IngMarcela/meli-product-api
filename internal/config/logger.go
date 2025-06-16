package config

import (
	"log"
	"os"
)

// Logger es el logger global de la aplicación
var Logger *log.Logger

// InitLogger inicializa el logger global
func InitLogger() {
	Logger = log.New(os.Stdout, "[meli-api] ", log.LstdFlags|log.Lshortfile)
}
