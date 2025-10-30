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
