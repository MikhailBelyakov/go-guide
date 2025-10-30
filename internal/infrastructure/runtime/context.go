package runtime

import (
	"arch/internal/config"
	mail2 "arch/internal/infrastructure/provider/mail"
	"context"
	"github.com/rs/zerolog"
	"log/slog"

	articleInfra "arch/internal/infrastructure/article"
	service "arch/internal/service/article"
	articleCreateCase "arch/internal/usecase/article/create"
)

type contextKey int

const (
	_ contextKey = iota
	loggerKey
	zerologKey
	mailProviderKey
	articleServiceKey
)

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func Logger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}

	panic("logger is not set")
}

func WithZerolog(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, zerologKey, logger)
}

func Zerolog(ctx context.Context) zerolog.Logger {
	if logger, ok := ctx.Value(zerologKey).(zerolog.Logger); ok {
		return logger
	}

	panic("zerolog is not set")
}

func WithMailProvider(ctx context.Context, provider *mail2.Provider) context.Context {
	return context.WithValue(ctx, mailProviderKey, provider)
}

func WithArticleService(ctx context.Context, service *service.AdminService) context.Context {
	return context.WithValue(ctx, articleServiceKey, service)
}

// ArticleService возвращает articleService из контекста
func ArticleService(ctx context.Context) *service.AdminService {
	if s, ok := ctx.Value(articleServiceKey).(*service.AdminService); ok {
		return s
	}
	panic("article service is not set in context")
}

// MailProvider возвращает mailProvider из контекста
func MailProvider(ctx context.Context) *mail2.Provider {
	if mp, ok := ctx.Value(mailProviderKey).(*mail2.Provider); ok {
		return mp
	}
	panic("mail provider is not set in context")
}

// BuildAppContext создаёт контекст с инициализацией всех зависимостей
func BuildAppContext(parent context.Context, config *config.Config) context.Context {
	ctx := parent

	mailClient := mail2.NewClient(config.MailAPIKey)
	mailProvider := mail2.NewProvider(*mailClient)
	ctx = WithMailProvider(ctx, mailProvider)

	repo := articleInfra.NewArticleRepo()

	articleUC := articleCreateCase.New(repo)

	articleService := service.NewAdminService(Zerolog(ctx), articleUC, mailProvider)
	ctx = WithArticleService(ctx, articleService)

	return ctx
}
