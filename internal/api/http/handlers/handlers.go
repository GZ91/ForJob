package handlers

import (
	"encoding/json"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"net/http"
	"regexp"
)

type handlers struct {
	nodeService handlerserService
	conf        *config.Config
	URLFilter   *regexp.Regexp
}

func New(nodeService handlerserService, conf *config.Config) *handlers {
	return &handlers{nodeService: nodeService, URLFilter: regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?(\w+\.[^:\/\n]+)`), conf: conf}
}

func (h *handlers) GetURLsToken(w http.ResponseWriter, r *http.Request) {
	var token string
	var tokenIDCTX models.CtxString = "token"
	TokenIDVal := r.Context().Value(tokenIDCTX)
	if TokenIDVal != nil {
		token = TokenIDVal.(string)
	}
	if token == "" {
		logger.Log.Info("trying to execute a method to retrieve a URL by a user by an unauthorized user")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	returnedURLs, err := h.nodeService.GetURLsToken(r.Context(), token)
	if err != nil {
		logger.Log.Error("when getting URLs on the user side", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(returnedURLs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	jsonText, err := json.Marshal(returnedURLs)
	if err != nil {
		logger.Log.Error("when creating a json file in the URL return procedure by user", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonText)
	if err != nil {
		logger.Log.Error("response recording error", zap.Error(err))
	}

}

func (h *handlers) AddBatchLinks(w http.ResponseWriter, r *http.Request) {
	StatusReturn := http.StatusCreated
	textBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var incomingBatchURL []string

	err = json.Unmarshal(textBody, &incomingBatchURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	releasedBatchURL, err := h.nodeService.AddBatchLink(r.Context(), incomingBatchURL)
	if err != nil {
		if errors.Is(err, errorsapp.ErrLinkAlreadyExists) {
			StatusReturn = http.StatusConflict
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}

	res, err := json.Marshal(releasedBatchURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if len(res) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(StatusReturn)
	_, err = w.Write(res)
	if err != nil {
		logger.Log.Error("response recording error", zap.String("error", err.Error()))
	}
}
