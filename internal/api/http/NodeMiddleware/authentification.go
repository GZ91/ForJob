package NodeMiddleware

import (
	"context"
	"fmt"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"net/http"
)

type NodeUse struct {
	confg *config.Config
}

func New(conf *config.Config) *NodeUse {
	return &NodeUse{confg: conf}
}

func (n *NodeUse) Authentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.String() == "/ping" {
			h.ServeHTTP(w, r)
			return
		}
		mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("Method", r.Method), zap.String("remote str", r.RemoteAddr)}
		token := ""
		cookie, err := r.Cookie("Authorization")
		if err != nil && err != http.ErrNoCookie {
			logger.Mserror("authorization", err, mainLog)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ok bool
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
			token, ok, err = validGetAuthentication(cookie.Value, n.confg.GetSecretKey())
			if err != nil {
				mainLog = append(mainLog, zap.String("secretkey", n.confg.GetSecretKey()))
				logger.Mserror("authorization", err, mainLog)
				return
			}
		}
		if ok {
			if n.confg.GetRootToken() == token {
				var tokenIDCTX models.CtxString = "token"
				r = r.WithContext(context.WithValue(r.Context(), tokenIDCTX, token))
				h.ServeHTTP(w, r)
				return
			}

		}
		logger.Msinfo("non authorization", nil, mainLog)
		w.WriteHeader(http.StatusUnauthorized)
		return
	})
}

func validGetAuthentication(tokenString, secret_key string) (string, bool, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				strErr := fmt.Sprintf("unexpected signing method: %v", t.Header["alg"])
				logger.Log.Error(strErr)
				return nil, fmt.Errorf(strErr)
			}
			return []byte(secret_key), nil
		})
	if err != nil {
		return "", false, err
	}

	if !token.Valid {
		return "", false, nil
	}

	return claims.Token, true, nil
}
