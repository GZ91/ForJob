package handlers

import (
	"github.com/GZ91/linkreduct/internal/app/config"
)

type handlers struct {
	nodeService HandlerserService
	conf        *config.Config
}

func New(nodeService HandlerserService, conf *config.Config) *handlers {
	return &handlers{nodeService: nodeService, conf: conf}
}
