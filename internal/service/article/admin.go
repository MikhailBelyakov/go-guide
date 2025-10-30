package article

import (
	"arch/internal/infrastructure/provider/mail"
	articleInterfaceType "arch/internal/interfaces/types"
	article "arch/internal/usecase/article/create"
	"context"
	"github.com/rs/zerolog"
)

type AdminService struct {
	logger       zerolog.Logger
	create       Create
	notification MailProvider
}

func NewAdminService(logger zerolog.Logger, create Create, notification MailProvider) *AdminService {
	return &AdminService{logger: logger, create: create, notification: notification}
}

func (s *AdminService) Create(ctx context.Context, userID uint64, _dto articleInterfaceType.RequestCreate) (articleInterfaceType.ResponseCreate, error) {
	model, err := s.create.Handler(ctx, article.CreateDTO{
		Title: _dto.Title,
	})
	if err != nil {
		return articleInterfaceType.ResponseCreate{}, err
	}

	err = s.notification.SendEmail(ctx, mail.SendDTO{})
	if err != nil {
		return articleInterfaceType.ResponseCreate{}, err
	}

	return articleInterfaceType.ResponseCreate{ID: model.ID}, nil
}
