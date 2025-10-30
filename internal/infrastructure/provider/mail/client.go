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
