package handlers

import (
	"context"
)

//go:generate mockery --name HandlerserService --with-expecter
type HandlerserService interface {
	GetTokens(ctx context.Context, NameService []string) (map[string]string, error)
	Ping(ctx context.Context) error
	GetURL(context.Context, string) (string, bool, error)
	GetServices(ctx context.Context, name string) (map[string]string, error)
	DeleteToken(ctx context.Context, token string) error
	GetLinks(ctx context.Context, token string) (map[string]string, error)
	DeleteLinkByLongLink(ctx context.Context, longLink string, token string) error
	DeleteLinkByShortLink(ctx context.Context, shortLink string, token string) error
	GetSmallLink(context.Context, string) (string, error)
}
