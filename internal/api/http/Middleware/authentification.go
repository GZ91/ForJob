package Middleware

import (
	"context"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/GZ91/linkreduct/internal/service"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (n *NodeUse) Authentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method), zap.String("remote str", r.RemoteAddr)}
		if service.CheckURL(r.URL.String()) {
			h.ServeHTTP(w, r)
			return
		}
		token := r.Header.Get("Authorization")
		ok := token != ""
		if ok {
			if n.confg.GetRootToken() == token {
				var tokenIDCTX models.CtxString = "Authorization"
				r = r.WithContext(context.WithValue(r.Context(), tokenIDCTX, token))
				h.ServeHTTP(w, r)
				return
			}
			okey, err := n.CheckToken(r.Context(), token)
			if err != nil {
				logger.Mserror("error when trying to check the presence of a token in the database", err, mainLog)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if okey {
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
