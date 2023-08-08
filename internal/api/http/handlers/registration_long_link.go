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

	textBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Log.Error("error read body", zap.Error(err))
		w.Write([]byte(err.Error()))
		return
	}

	var data models.RequestData

	err = json.Unmarshal(textBody, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Log.Error("error in unmarshal", zap.Error(err))
		w.Write([]byte(err.Error()))
		return
	}

	link := data.URL

	if !h.URLFilter.MatchString(link) {
		w.WriteHeader(http.StatusBadRequest)
		logger.Log.Error("string is not a reference")
		return
	}

	bodyText, err := h.nodeService.GetSmallLink(r.Context(), link)
	if err != nil {
		if errors.Is(err, errorsapp.ErrLinkAlreadyExists) {
			StrData := &models.ReturnData{LongLink: link, ShortLink: bodyText}
			Data, err := json.Marshal(StrData)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				logger.Log.Error("error in marshal", zap.Error(err))
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
			logger.Log.Error("error in work DB", zap.Error(err))
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
		logger.Log.Error("error marshal", zap.Error(err))
		return
	}
	if len(res) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		logger.Log.Error("non record", zap.Error(err))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		logger.Log.Error("response recording error", zap.Error(err))
	}
}
