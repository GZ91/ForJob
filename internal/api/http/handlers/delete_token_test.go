package handlers

import (
	"context"
	"errors"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/GZ91/linkreduct/internal/service"
	"github.com/GZ91/linkreduct/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/go-chi/chi/v5"
	mock_test "github.com/stretchr/testify/mock"
)

func TestDeleteToken(t *testing.T) {
	logger.Initializing("info")
	mockStorager := &mocks.Storeger{}

	configa := config.New(false, "", "", 100, 4, "", "existent_token")
	serviceNode := service.New(context.Background(), mockStorager, configa)

	h := &handlers{
		conf:        configa,
		nodeService: serviceNode,
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

	mockStorager.On("DeleteToken", mock_test.Anything, "existent_token").Return(nil)
	mockStorager.On("DeleteToken", mock_test.Anything, "not_found_token").Return(errorsapp.ErrNotFoundToken)
	mockStorager.On("DeleteToken", mock_test.Anything, "other_token").Return(errors.New("other"))

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
