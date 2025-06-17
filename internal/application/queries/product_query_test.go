package queries_test

import (
	"testing"

	"meli-product-api/internal/application/queries"
	"meli-product-api/internal/domain/model"

	"github.com/stretchr/testify/assert"
)

type mockRepo struct{}

func (m *mockRepo) GetByID(id string) (*model.Product, error) {
	return &model.Product{ID: id, Title: "Producto Test"}, nil
}

func TestGetProductByIDQuery_Execute(t *testing.T) {
	q := queries.NewGetProductByIDQuery(&mockRepo{})

	result, err := q.Execute("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Producto Test", result.Title)
}
