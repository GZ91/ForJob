package service

import (
	"context"
	"github.com/GZ91/linkreduct/internal/models"
)

// Storeger
//
//go:generate mockery --name Storeger --with-expecter
type Storeger interface {
	AddURL(context.Context, string) (string, error)
	GetURL(context.Context, string) (string, bool, error)
	Ping(context.Context) error
	GetTokens(ctx context.Context, namesServices []string) (map[string]string, error)
	AddBatchLink(context.Context, []string) (map[string]string, error)
	FindLongURL(context.Context, string) (string, bool, error)
	GetLinksToken(context.Context, string) ([]models.ReturnedStructURL, error)
	InitializingRemovalChannel(context.Context, chan []models.StructDelURLs) error
}

// Storeger
//
//go:generate mockery --name ConfigerService --with-expecter
type ConfigerService interface {
	GetAddressServerURL() string
	GetSecretKey() string
}
