package products

import (
	"context"
	repo "github.com/Violetory/e-com/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(context context.Context) ([]repo.Product, error)
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{repo: repo}
}

func (s *service) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}
