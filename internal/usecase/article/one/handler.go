package article

import (
	infra "arch/internal/infrastructure/article"
	"context"
)

type UseCase struct {
	repository Repository
}

func New(repository Repository) *UseCase {
	return &UseCase{repository: repository}
}

func (uc *UseCase) Handler(ctx context.Context, id string) (*infra.Entity, error) {
	return uc.repository.GetByID(ctx, id)
}
