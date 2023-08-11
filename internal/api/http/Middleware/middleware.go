package Middleware

import (
	"context"
	"github.com/GZ91/linkreduct/internal/app/config"
)

type NodeUse struct {
	confg *config.Config
	servces
}

type servces interface {
	CheckToken(ctx context.Context, Token string) (bool, error)
}

func New(conf *config.Config, servces servces) *NodeUse {
	return &NodeUse{confg: conf, servces: servces}
}
