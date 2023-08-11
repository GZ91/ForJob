package handlers

import (
	"encoding/json"
	"errors"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *handlers) GetToken(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "token"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}
	if token != h.conf.GetRootToken() {
		logger.Msinfo("not enough rights to issue the token", nil, nil)
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		w.Write([]byte("not enough rights to issue the token"))
		return
	}

	textBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Msinfo("I couldn't read the body", nil, nil)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	mainLog := []zap.Field{zap.String("token", token), zap.ByteString("body", textBody), zap.String("URL", r.URL.String()), zap.String("Method", r.Method)}
	if err != nil {
		logger.Mserror("when creating a new token", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var serviceSTR models.ServiceNamesIn
	err = json.Unmarshal(textBody, &serviceSTR)
	if err != nil {
		logger.Mserror("when translating data from json format to structure", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := h.nodeService.GetTokens(r.Context(), serviceSTR.Services)
	if err != nil {
		if errors.Is(err, errorsapp.ErrAlredyBeenRegistered) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if len(data) == 0 {
		logger.Mserror("token was not created, error on lower layers", nil, mainLog)
		return
	}
	w.Header().Add(" Content-Type", "application/json")
	dataout, err := json.Marshal(data)
	if err != nil {
		logger.Mserror("when converting a list of key:value pairs", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	_, err = w.Write(dataout)
	if err != nil {
		logger.Mserror("json records in the response body", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
