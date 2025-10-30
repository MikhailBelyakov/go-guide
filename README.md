**Гайдлайн по архитектуре Go-проекта**


---

## 1. Пример структуры проекта

```
cmd
├─ build.go
├─ init.go
├─ root.go
├─ serve.go
├─ serve_private.go
└─ serve_public.go
internal
├─ config
│  └─ config.go
├─ domain
│  └─ article
│     ├─ constants.go
│     └─ entity.go
├─ event
│  └─ article
│     └─ created.go
├─ infrastructure
│  ├─ article
│  │  ├─ const.go
│  │  ├─ entity.go
│  │  ├─ factory.go
│  │  └─ repository.go
│  ├─ provider
│  │  └─ mail
│  │     ├─ client.go
│  │     ├─ dto.go
│  │     └─ provider.go
│  └─ runtime
│     └─ context.go
├─ interfaces
│  ├─ types
│  ├─ article_v1.go
│  └─ tg.go
├─ service
│  └─ article
│     ├─ admin.go
│     ├─ public.go
│     └─ dependency.go
└─ transport
usecase
└─ article
   ├─ create
   │  ├─ dependency.go
   │  ├─ handler.go
   │  └─ validator.go
   └─ one
      ├─ dependency.go
      ├─ handler.go
      └─ validator.go
pkg

```

---

## 2. Описание слоёв

### 2.1 cmd
- Точка входа в приложение.
- Используется Cobra для CLI-команд (`serve`, `migrate`, `worker`).

### 2.2 internal/config
- Конфигурации приложения.
- Загружаются через env или конфиг-файлы.

### 2.3 internal/domain (опциональный слой на сложных проектах)
- Бизнес-сущности и константы.
- Каждая сущность в отдельной папке (`article/`):
    - `entity.go` — структура сущности
    - `constants.go` — статусы, типы, ограничения

### 2.4 internal/event
- Доменные события (`ArticleCreated`) в отдельных папках для каждой сущности.
- Используются для реакции внутри usecase.

### 2.5 internal/infrastructure
- Работа с внешними ресурсами и системами:
    - article/ — репозитории, фабрики, модели
    - provider/ — внешние интеграции (Mail, Telegram)
        - client.go — подключение
        - dto.go — обмен данными между провайдером и проектом
        - provider.go — функции для использования клиента
    - runtime/ — контексты, логгеры, утилиты

### 2.6 internal/interfaces
- Контракты внешних API.
- Общие типы данных в `types/`.

### 2.7 internal/service
- Адаптация usecase под внешние интерфейсы.
- Зависимости для сервисов хранятся рядом (`dependency.go`).
- Обрабатывает DTO из транспорта и вызывает usecase/repository.
- Примеры сервисов на одну сущность:
    - `admin.go` — административные операции
    - `public.go` — публичные операции

### 2.8 internal/transport
- Генерация роутеров, middleware, логирование, метрики.
- Не содержит бизнес-логику.

### 2.9 internal/usecase
- Бизнес-логика.
- Каждая операция (create, update, one) в отдельной папке.
- Зависимости для usecase хранятся рядом (`dependency.go`).

### 2.10 pkg
- Общие пакеты и утилиты, не зависящие от конкретной сущности.

---

## 3. DTO и константы

- Usecase DTO — internal/usecase/<entity>/<operation>/handler.go
- Provider DTO — internal/infrastructure/provider/<name>/dto.go
- Константы, типы для сущности — domain/<entity>/const.go, domain/<entity>/type.go  (если есть слой domain)
- Константы, типы  для  сущности — infrastructure/<entity>/const.go , infrastructure/<entity>/type.go (если есть слой domain)

---

## 4. События

- Доменные события → internal/event/<entity>/

---

## 5. Наименование

- Пакеты: маленькими буквами (`article`, `dto`)
- Файлы: отражают содержимое (`repository.go`, `dto.go`, `service.go`)

---

## 6. Взаимодействие слоёв

```
transport -> service -> usecase -> domain(опционально) -> infrastructure
                      ↑                    ↓
          DTO / события               репозитории, провайдеры
```

---

## 7. Стандарты кода

1. Все слои тестируются отдельно
2. Usecase покрываются unit-тестами
3. Интеграционные тесты для service/infrastructure
4. Контекст (`context.Context`) передаётся через все слои
5. DI через runtime
6. Ошибки через `error`, не panic

---

## 8. Примеры содержимого файлов

**Service (admin.go)**
```go
package article

import (
	infra "arch/internal/infrastructure/article"
	"arch/internal/infrastructure/provider/mail"
	articleInterfaceType "arch/internal/interfaces/types"
	article "arch/internal/usecase/article/create"
	"context"
	"github.com/rs/zerolog"
)

type Create interface {
	Handler(ctx context.Context, dto article.CreateDTO) (*infra.Entity, error)
}

type MailProvider interface {
	SendEmail(ctx context.Context, dto mail.SendDTO) error
}

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
```

**Provider (mail/provider.go)**
```go
package mail

import "context"

type Provider struct {
	client Client
}

func NewProvider(client Client) *Provider {
	return &Provider{client: client}
}

func (p *Provider) SendEmail(ctx context.Context, dto SendDTO) error {
	return p.client.Request(ctx, dto.To, dto.Subject, dto.Body)
}
```

**Provider DTO (mail/dto.go)**
```go
package mail

type SendDTO struct {
	To      string
	Subject string
	Body    string
}
```

**Provider Client (mail/client.go)**
```go
package mail

import "context"

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) Request(ctx context.Context, to, subject, body string) error {
	return nil
}
```

**UseCase Handler (usecase/article/create/handler.go)**
```go
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
	validator Validator
}

func New(repository Repository, validator Validator) *UseCase {
	return &UseCase{repository: repository, validator: validator}
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
```

**UseCase Dependency (usecase/article/create/dependency.go)**
```go
package article

import (
	"arch/internal/infrastructure/article"
	"context"
)

type Repository interface {
	SaveArticle(ctx context.Context, model article.Entity) error
}


type Validator interface {
	Validate(ctx context.Context, dto CreateDTO) error
}

```

**UseCase Validator (usecase/article/create/validator.go)**
```go
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

```

