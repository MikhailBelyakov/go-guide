package article

import (
	articleInfra "arch/internal/infrastructure/article"
	"arch/internal/infrastructure/provider/mail"
	article "arch/internal/usecase/article/create"
	"context"
)

type Create interface {
	Handler(ctx context.Context, dto article.CreateDTO) (*articleInfra.Entity, error)
}

type MailProvider interface {
	SendEmail(ctx context.Context, dto mail.SendDTO) error
}

type One interface {
	Handler(ctx context.Context, id string) (*articleInfra.Entity, error)
}
