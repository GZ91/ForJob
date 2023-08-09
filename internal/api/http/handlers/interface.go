package handlers

import (
	"context"
	"github.com/GZ91/linkreduct/internal/models"
)

//go:generate mockery --name handlerserService --with-expecter
type handlerserService interface {
	GetNewTokens(ctx context.Context, NameService []string) (map[string]string, error)
	Ping(ctx context.Context) error
	GetURL(context.Context, string) (string, bool, error)

	GetSmallLink(context.Context, string) (string, error)
	GetURLsToken(context.Context, string) ([]models.ReturnedStructURL, error)
	AddBatchLink(context.Context, []string) (map[string]string, error)
	DeletedLinks([]string, string)
}
