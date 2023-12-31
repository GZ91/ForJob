package service

import (
	"context"
)

// Storeger
//
//go:generate mockery --name Storeger --with-expecter
type Storeger interface {
	AddURL(context.Context, string) (string, error)
	GetURL(context.Context, string) (string, bool, error)
	Ping(context.Context) error
	GetTokens(ctx context.Context, namesServices []string) (map[string]string, error)
	CheckToken(ctx context.Context, token string) (bool, error)
	DeleteToken(ctx context.Context, token string) error
	FindLongURL(context.Context, string, string) (string, bool, error)
	GetServices(ctx context.Context, name string) (map[string]string, error)
	GetLinks(ctx context.Context, token string) (map[string]string, error)
	DeleteLinkByLongLink(ctx context.Context, longLink string, token string) error
	DeleteLinkByShortLink(ctx context.Context, shortLink string, token string) error
}

// Configer
//
//go:generate mockery --name ConfigerService --with-expecter
type ConfigerService interface {
	GetAddressServerURL() string
	GetSecretKey() string
}
