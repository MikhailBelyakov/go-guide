package article

import (
	articleHttpType "arch/internal/interfaces/types"
	"context"
)

type PublicService struct {
	oneUseCase One
}

func NewPublicService(oneUseCase One) *PublicService {
	return &PublicService{oneUseCase: oneUseCase}
}

func (s *PublicService) One(ctx context.Context, id string) (*articleHttpType.ResponseOne, error) {
	result, err := s.oneUseCase.Handler(ctx, id)
	if err != nil {
		return nil, err
	}

	return &articleHttpType.ResponseOne{Title: result.Title}, nil
}
