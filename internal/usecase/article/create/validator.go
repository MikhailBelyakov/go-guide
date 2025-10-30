package article

import (
	"context"
	"errors"
)

type validator struct {
}

func NewValidator() Validator {
	return &validator{}
}

func (v *validator) Validate(_ context.Context, dto CreateDTO) error {
	if dto.Title == "" {
		return errors.New("title cannot be empty")
	}

	return nil
}
