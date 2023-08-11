package handlers

import (
	"encoding/json"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"go.uber.org/zap"
	"net/http"
)

func (h *handlers) GetServices(w http.ResponseWriter, r *http.Request) {
	mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method)}
	nameService := r.Header.Get("name")
	servces, err := h.nodeService.GetServices(r.Context(), nameService)
	if err != nil {
		logger.Mserror("when the service is retrieved from the database", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(servces)
	if err != nil {
		logger.Mserror("when encoding into json format", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if len(data) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
