package handlers

import (
	"encoding/json"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *handlers) GetServices(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "token"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}
	if token != h.conf.GetRootToken() {
		logger.Msinfo("insufficient rights to view services", nil, nil)
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		w.Write([]byte("insufficient rights to view services"))
		return
	}

	mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method)}
	dataName, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Msinfo("in reading the body", err, mainLog)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body reading error"))
	}
	type Name struct {
		Name string `json:"name"`
	}
	var name Name
	json.Unmarshal(dataName, &name)
	services, err := h.nodeService.GetServices(r.Context(), name.Name)
	if err != nil {
		logger.Mserror("when the service is retrieved from the database", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(services)
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
