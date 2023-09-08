package handlers

import (
	"errors"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"go.uber.org/zap"
	"net/http"
)

func (h *handlers) DeleteLinkByLongLink(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "Authorization"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}

	longLink := r.URL.Query().Get("url")

	mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("longLink", longLink), zap.String("token", token)}

	err := h.nodeService.DeleteLinkByLongLink(r.Context(), longLink, token)
	if err != nil {
		if errors.Is(err, errorsapp.ErrNotFoundLink) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		logger.Mserror("link removal error", err, mainLog)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
