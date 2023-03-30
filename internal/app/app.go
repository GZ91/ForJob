package app

import (
	"fmt"
	"github.com/GZ91/linkreduct/internal/config"
	"github.com/GZ91/linkreduct/internal/server"
)

var appLink *app

type app struct {
	Config config.Config
}

func New(config config.Config) *app {
	if appLink == nil {
		appLink = &app{
			config,
		}
		return appLink
	}
	return appLink
}

func (r app) Run() {
	if err := server.Start(r.Config); err != nil {
		fmt.Printf("%v \n", err)
	}
}
