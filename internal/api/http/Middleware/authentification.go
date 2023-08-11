package Middleware

import (
	"context"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (n *NodeUse) Authentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			h.ServeHTTP(w, r)
			return
		}
		mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method), zap.String("remote str", r.RemoteAddr)}
		token := r.Header.Get("Authorization")
		ok := token != ""
		if ok {
			if n.confg.GetRootToken() == token {
				var tokenIDCTX models.CtxString = "token"
				r = r.WithContext(context.WithValue(r.Context(), tokenIDCTX, token))
				h.ServeHTTP(w, r)
				return
			}
			if n.CheckToken(token) {
				if r.URL.String() != "/token" && !strings.HasPrefix(r.URL.String(), "/services") {
					h.ServeHTTP(w, r)
					return
				}
			}
		}
		logger.Msinfo("non authorization", nil, mainLog)
		w.WriteHeader(http.StatusUnauthorized)
		return
	})
}
