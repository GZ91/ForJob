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

func (h *handlers) AddLongLinkJSON(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "Authorization"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}
	mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method), zap.String("remote str", r.RemoteAddr), zap.String("token", token)}
	textBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Mserror("error read body", err, mainLog)
		w.Write([]byte(err.Error()))
		return
	}
	mainLog = append(mainLog, zap.ByteString("body", textBody))

	var data models.RequestData

	err = json.Unmarshal(textBody, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Mserror("error in unmarshal", err, mainLog)
		w.Write([]byte(err.Error()))
		return
	}

	link := data.URL

	bodyText, err := h.nodeService.GetSmallLink(r.Context(), link, token)
	if err != nil {
		if errors.Is(err, errorsapp.ErrLinkAlreadyExists) {
			StrData := &models.ReturnData{LongLink: link, ShortLink: bodyText}
			Data, err := json.Marshal(StrData)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				logger.Mserror("error in marshal", err, mainLog)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(Data)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			logger.Mserror("error in work DB", err, mainLog)
			return
		}
	}
	if bodyText == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		logger.Log.Info("in body not record")
		return
	}

	Data := models.ReturnData{LongLink: link, ShortLink: bodyText}

	res, err := json.Marshal(Data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		logger.Mserror("error marshal", err, mainLog)
		return
	}
	if len(res) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		logger.Mserror("non record", err, mainLog)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		logger.Mserror("response recording error", err, mainLog)
	}
}
