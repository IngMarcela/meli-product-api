package usecases

import (
	"strings"

	"meli-product-api/internal/domain/model"
	"meli-product-api/internal/domain/ports"
)

type SearchProductsUseCase struct {
	Repo ports.ProductRepository
}

func NewSearchProductsUseCase(repo ports.ProductRepository) *SearchProductsUseCase {
	return &SearchProductsUseCase{Repo: repo}
}

func (uc *SearchProductsUseCase) Execute(query string) ([]model.Product, error) {
	all, err := uc.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []model.Product
	q := strings.ToLower(query)

	for _, p := range all {
		if strings.Contains(strings.ToLower(p.Title), q) || strings.Contains(strings.ToLower(p.Description), q) {
			result = append(result, p)
		}
	}

	return result, nil
}
