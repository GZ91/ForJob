package handlers

import (
	"github.com/GZ91/linkreduct/internal/app/logger"
	"go.uber.org/zap"
	"net/http"
)

func (h *handlers) PingDataBase(w http.ResponseWriter, r *http.Request) {
	mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method)}
	err := h.nodeService.Ping(r.Context())
	if err != nil {
		logger.Mserror("ping error", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
