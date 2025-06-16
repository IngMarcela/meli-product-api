package queries_test

import (
	"testing"

	"meli-product-api/internal/application/queries"
	"meli-product-api/internal/domain/model"

	"github.com/stretchr/testify/assert"
)

type mockRepo struct{}

func (m *mockRepo) GetAll() ([]model.Product, error) {
	return []model.Product{{ID: "1", Title: "Producto Test"}}, nil
}

func (m *mockRepo) GetByID(id string) (*model.Product, error) {
	return &model.Product{ID: id, Title: "Producto Test"}, nil
}

func TestGetAllProductsQuery_Execute(t *testing.T) {
	q := queries.NewGetAllProductsQuery(&mockRepo{})

	result, err := q.Execute()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].ID)
}

func TestGetProductByIDQuery_Execute(t *testing.T) {
	q := queries.NewGetProductByIDQuery(&mockRepo{})

	result, err := q.Execute("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Producto Test", result.Title)
}
