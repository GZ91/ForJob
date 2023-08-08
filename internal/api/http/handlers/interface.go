package handlers

import (
	"context"
	"github.com/GZ91/linkreduct/internal/models"
)

//go:generate mockery --name handlerserService --with-expecter
type handlerserService interface {
	GetSmallLink(context.Context, string) (string, error)
	GetURL(context.Context, string) (string, bool, error)
	GetURLsToken(context.Context, string) ([]models.ReturnedStructURL, error)
	Ping(ctx context.Context) error
	AddBatchLink(context.Context, []string) (map[string]string, error)
	DeletedLinks([]string, string)
}
