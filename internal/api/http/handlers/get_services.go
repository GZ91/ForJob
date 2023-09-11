package handlers

import (
	"encoding/json"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"go.uber.org/zap"
	"net/http"
)

func (h *handlers) GetServices(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "Authorization"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}
	if token != h.conf.GetRootToken() {
		logger.Msinfo("insufficient rights to view services", nil, nil)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("insufficient rights to view services"))
		return
	}

	mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method)}

	Name := r.URL.Query().Get("name")

	services, err := h.nodeService.GetServices(r.Context(), Name)
	if err != nil {
		logger.Mserror("when the service is retrieved from the database", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(services) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	data, err := json.Marshal(services)
	if err != nil {
		logger.Mserror("when encoding into json format", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
