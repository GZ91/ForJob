package handlers

import (
	"encoding/json"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"go.uber.org/zap"
	"net/http"
)

type dataReturnLinks struct {
	Links []models.ReturnData `json:"links"`
}

func (h *handlers) GetLinks(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "Authorization"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}
	mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method), zap.String("token", token)}

	data, err := h.nodeService.GetLinks(r.Context(), token)
	if err != nil {
		logger.Mserror("link retrieval error", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(data) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var dataReturn dataReturnLinks

	for shortLink, longLink := range data {
		dataReturn.Links = append(dataReturn.Links, models.ReturnData{LongLink: longLink, ShortLink: shortLink})
	}
	dataJson, err := json.Marshal(dataReturn)
	if err != nil {
		logger.Mserror("when converting to json", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(dataJson)
	if err != nil {
		logger.Mserror("when recording a message", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
