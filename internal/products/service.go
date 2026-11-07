package products

import (
	"context"

	repo "github.com/nnaxim/go-backend-study/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product ,error)
	FindProductByID(ctx context.Context, id int64) (repo.Product, error)
}

type srv struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &srv{
		repo: repo,
	}
}

func (s *srv) ListProducts(ctx context.Context) ([]repo.Product ,error) {
	return s.repo.ListProducts(ctx)
}

func (s *srv) FindProductByID(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductByID(ctx, id)
}