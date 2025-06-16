package config_test

import (
	"os"
	"testing"

	"meli-product-api/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_DefaultValues(t *testing.T) {
	os.Clearenv() // asegúrate de que no hay variables seteadas

	cfg := config.LoadConfig()

	assert.Equal(t, "3000", cfg.Port)
	assert.Equal(t, "data/products.json", cfg.ProductFile)
}
