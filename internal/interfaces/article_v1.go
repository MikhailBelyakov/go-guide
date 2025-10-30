package interfaces

import (
	article "arch/internal/interfaces/types"
	"context"
)

type ArticleV1 interface {
	Create(ctx context.Context, userID uint64, data article.RequestCreate) (response article.ResponseCreate, err error)
}
