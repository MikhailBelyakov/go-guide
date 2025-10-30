package article

import (
	"arch/internal/infrastructure/article"
	"context"
)

type Repository interface {
	SaveArticle(ctx context.Context, model article.Entity) error
}
