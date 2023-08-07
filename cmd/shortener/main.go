package main

import (
	"context"
	"github.com/GZ91/linkreduct/internal/app"
	"github.com/GZ91/linkreduct/internal/app/initializing"
)

func main() {
	ctx := context.Background()
	conf, err := initializing.Configuration()
	if err != nil {
		panic(err)
	}
	appliction := app.New(conf)
	appliction.Run(ctx)

}
