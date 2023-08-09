package server

import (
	"context"
	"github.com/GZ91/linkreduct/internal/api/http/NodeMiddleware"
	"github.com/GZ91/linkreduct/internal/api/http/handlers"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/app/signalreception"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/GZ91/linkreduct/internal/service"
	"github.com/GZ91/linkreduct/internal/service/genrunes"
	"github.com/GZ91/linkreduct/internal/storage/postgresql"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

type NodeStorager interface {
	service.Storeger
	Close() error
}

func Start(ctx context.Context, conf *config.Config) (er error) {
	var NodeStorage NodeStorager
	GeneratorRunes := genrunes.New()
	if conf.GetConfDB().Empty() {
		return errorsapp.ErrNotConfiguration
	}
	NodeStorage, err := postgresql.New(ctx, conf, GeneratorRunes)
	if err != nil {
		return err
	}

	NodeService := service.New(ctx, NodeStorage, conf, make(chan []models.StructDelURLs))
	handls := handlers.New(NodeService)

	router := chi.NewRouter()

	NodeUse := NodeMiddleware.New(conf)

	router.Use(NodeUse.Authentication)
	router.Use(NodeUse.WithLogging)
	router.Use(NodeUse.Compress)
	router.Use(NodeUse.CalculateSize)

	//
	//router.Get("/api/token/urls", handls.GetURLsToken)
	//	router.Post("/api/shorten/batch", handls.AddBatchLinks)

	//	router.Delete("/api/user/urls", handls.DeleteURLs)

	router.Get("/{id}", handls.GetLongURL)
	router.Get("/ping", handls.PingDataBase)
	router.Post("/token", handls.AddLongLinkJSON)

	Server := http.Server{}
	Server.Addr = conf.GetAddressServer()
	Server.Handler = router

	wg := sync.WaitGroup{}

	go signalreception.OnClose([]signalreception.Closer{
		&signalreception.Stopper{CloserInterf: &Server, Name: "server"},
		&signalreception.Stopper{CloserInterf: NodeStorage, Name: "node storage"},
		&signalreception.Stopper{CloserInterf: NodeService, Name: "node service"}},
		&wg)

	if err := Server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Error("server startup error", zap.String("error", err.Error()))
		}
	}
	wg.Wait()
	return

}
