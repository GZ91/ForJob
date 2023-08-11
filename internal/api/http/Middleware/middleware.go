package Middleware

import "github.com/GZ91/linkreduct/internal/app/config"

type NodeUse struct {
	confg *config.Config
	servces
}

type servces interface {
	CheckToken(Token string) bool
}

func New(conf *config.Config, servces servces) *NodeUse {
	return &NodeUse{confg: conf, servces: servces}
}
