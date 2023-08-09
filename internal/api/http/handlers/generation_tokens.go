package handlers

import (
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"net/http"
)

func (h *handlers) GetNewToken(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "token"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}
	if token != h.conf.GetRootToken() {
		logger.Msinfo("not enough rights to issue the token", nil, nil)
		return
	}
}
