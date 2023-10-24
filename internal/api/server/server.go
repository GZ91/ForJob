package server

import (
	"github.com/GZ91/EmploymentTest/internal/api/handlers"
	"github.com/GZ91/EmploymentTest/internal/app/logger"
	"github.com/GZ91/EmploymentTest/internal/service"
	"github.com/GZ91/EmploymentTest/internal/storage"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func Run(path string) error {
	NodeStorage, err := storage.New(path)
	if err != nil {
		logger.Log.Error("При создании слоя хранилища", zap.Error(err))
		return err
	}
	NodeService, err := service.New(NodeStorage)
	if err != nil {
		logger.Log.Error("При создании слоя сервиса", zap.Error(err))
		return err
	}
	NodeHandls, err := handlers.New(NodeService)
	if err != nil {
		logger.Log.Error("При создании слоя хэндлеров", zap.Error(err))
		return err
	}
	router := chi.NewRouter()

	router.Get("/get-items/{id}", NodeHandls.GetItems)

	Server := http.Server{}
	Server.Addr = ":8181"
	Server.Handler = router

	return Server.ListenAndServe()
}
