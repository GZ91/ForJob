package handlers

import (
	"context"
	"errors"
	"github.com/GZ91/linkreduct/internal/api/http/handlers/mocks"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/go-chi/chi/v5"
)

// mockNodeService - мок сервиса, реализующего метод DeleteToken
type mockNodeService struct {
	mocks.HandlerserService
}

func (m *mockNodeService) DeleteToken(ctx context.Context, token string) error {
	if token == "existent_token" {
		return nil
	} else if token == "not_found_token" {
		return errorsapp.ErrNotFoundToken
	}
	return errors.New("other_error")
}

func TestDeleteToken(t *testing.T) {
	logger.Initializing("info")
	mock := &mockNodeService{}
	h := &handlers{
		conf:        config.New(false, "", "", 100, 4, "", "existent_token"), // замените на вашу реальную структуру конфигурации
		nodeService: mock,
	}
	r := chi.NewRouter()
	r.HandleFunc("/delete/{token}", h.DeleteToken)

	testCases := []struct {
		name         string
		token        string
		expectedCode int
	}{
		{"AuthorizedDelete", "existent_token", http.StatusAccepted},
		{"NotFoundToken", "not_found_token", http.StatusNotFound},
		{"OtherError", "other_token", http.StatusInternalServerError},
	}
	mock.On("DeleteToken", "existent_token").Return(nil)
	mock.On("DeleteToken", "not_found_token").Return(errorsapp.ErrNotFoundToken)
	mock.On("DeleteToken", "other_token").Return(errors.New("other"))

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/delete/"+tc.token, nil)
			if err != nil {
				t.Fatal(err)
			}

			var tokenIDCTX models.CtxString = "Authorization"
			req = req.WithContext(context.WithValue(req.Context(), tokenIDCTX, "existent_token"))

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if rr.Code != tc.expectedCode {
				t.Errorf("expected status %d, got %d", tc.expectedCode, rr.Code)
			}
		})
	}
}
