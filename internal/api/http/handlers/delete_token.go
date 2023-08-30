package handlers

import (
	"errors"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func (h *handlers) DeleteToken(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "Authorization"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}
	if token != h.conf.GetRootToken() {
		logger.Msinfo("insufficient rights to delete the token", nil, nil)
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		w.Write([]byte("insufficient rights to delete the token"))
		return
	}
	mainLog := []zap.Field{zap.String("token", token), zap.String("URL", r.URL.String()), zap.String("Method", r.Method)}
	tokenDel := chi.URLParam(r, "token")

	if tokenDel == "" {
		logger.Msinfo("no deletion token specified", nil, mainLog)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.nodeService.DeleteToken(r.Context(), tokenDel)
	if err != nil {
		if errors.Is(err, errorsapp.ErrNotFoundToken) {
			logger.Msinfo("token not found", err, mainLog)
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			logger.Mserror("when the token is deleted", err, mainLog)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusAccepted)
}
