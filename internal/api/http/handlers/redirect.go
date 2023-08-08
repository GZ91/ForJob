package handlers

import (
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func (h *handlers) GetLongURL(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	mainLog := []zap.Field{zap.String("URL", r.URL.String()), zap.String("ID", id)}
	link, ok, err := h.nodeService.GetURL(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Mserror("error get URL", err, mainLog)
		return
	}
	if ok {
		w.Header().Add("Location", link)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
