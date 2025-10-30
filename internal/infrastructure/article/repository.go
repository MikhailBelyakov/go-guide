package article

import "context"

type Repository struct {
}

func NewArticleRepo() *Repository {
	return &Repository{}
}

func (r *Repository) Handler(ctx context.Context, model Entity) error {
	return nil
}
