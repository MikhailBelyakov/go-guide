package article

import (
	"arch/internal/infrastructure/article"
	"context"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (*article.Entity, error)
}
