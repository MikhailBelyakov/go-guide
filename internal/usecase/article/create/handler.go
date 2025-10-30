package article

import (
	infra "arch/internal/infrastructure/article"
	"context"
	"errors"
	"github.com/google/uuid"
)

type CreateDTO struct {
	Title string
}

type UseCase struct {
	repository Repository
}

func New(repository Repository) *UseCase {
	return &UseCase{repository: repository}
}

func (uc *UseCase) Handler(ctx context.Context, dto CreateDTO) (*infra.Entity, error) {
	if dto.Title == "" {
		return nil, errors.New("title cannot be empty")
	}

	article := infra.NewArticleModel(
		uuid.New().String(),
		dto.Title,
	)

	if err := uc.repository.SaveArticle(ctx, *article); err != nil {
		return nil, err
	}

	return article, nil
}
