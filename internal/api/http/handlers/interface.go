package handlers

import (
	"context"
	"github.com/GZ91/linkreduct/internal/models"
)

//go:generate mockery --name handlerserService --with-expecter
type handlerserService interface {
	GetTokens(ctx context.Context, NameService []string) (map[string]string, error)
	Ping(ctx context.Context) error
	GetURL(context.Context, string) (string, bool, error)
	GetServices(ctx context.Context, name string) (map[string]string, error)
	DeleteToken(ctx context.Context, token string) error

	GetSmallLink(context.Context, string) (string, error)
	GetURLsToken(context.Context, string) ([]models.ReturnedStructURL, error)
	AddBatchLink(context.Context, []string) (map[string]string, error)
}
