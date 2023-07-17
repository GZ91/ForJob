package handlers

import (
	"github.com/GZ91/EmploymentTest/internal/app/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Service interface {
	RevertSearchStructures(id string) ([]byte, error)
}

type Handlers struct {
	NodeService Service
}

func New(service Service) (*Handlers, error) {
	return &Handlers{NodeService: service}, nil
}

func (h Handlers) GetItems(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := h.NodeService.RevertSearchStructures(id)
	if err != nil {
		logger.Log.Error("Ошибка при получении данных по отбору")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		logger.Log.Error("Ошибка при записи сообщения")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
