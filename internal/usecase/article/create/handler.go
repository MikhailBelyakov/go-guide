package article

import (
	infra "arch/internal/infrastructure/article"
	"context"
	"github.com/google/uuid"
)

type CreateDTO struct {
	Title string
}

type UseCase struct {
	repository Repository
	validator  Validator
}

func New(repository Repository) *UseCase {
	return &UseCase{repository: repository}
}

func (uc *UseCase) Handler(ctx context.Context, dto CreateDTO) (*infra.Entity, error) {
	if err := uc.validator.Validate(ctx, dto); err != nil {
		return nil, err
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
