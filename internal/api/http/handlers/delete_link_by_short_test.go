package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/GZ91/linkreduct/internal/service"
	"github.com/GZ91/linkreduct/internal/service/mocks"
	"github.com/go-chi/chi/v5"
	mock_test "github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_handlers_DeleteLinkByShortLink(t *testing.T) {
	logger.Initializing("info")

	mockStorager := &mocks.Storeger{}
	configa := config.New(false, "", "", 100, 4, "", "existent_token")

	serviceNode := service.New(context.Background(), mockStorager, configa)

	h := &handlers{
		conf:        configa, // замените на вашу реальную структуру конфигурации
		nodeService: serviceNode,
	}

	r := chi.NewRouter()
	r.HandleFunc("/links/short", h.DeleteLinkByShortLink)

	type TestCase struct {
		name         string
		token        string
		param        string
		expectedCode int
		funcMock     func()
	}

	testCases := []TestCase{
		{
			name:         "other error",
			token:        "existent_token1",
			param:        `shortURL`,
			expectedCode: http.StatusInternalServerError,
			funcMock: func() {
				mockStorager.On("DeleteLinkByShortLink", mock_test.Anything, "shortURL", "existent_token1").Return(errors.New("other"))
			},
		},
		{
			name:         "status ok",
			token:        "existent_token2",
			param:        `shortURL`,
			expectedCode: http.StatusOK,
			funcMock: func() {
				mockStorager.On("DeleteLinkByShortLink", mock_test.Anything, "shortURL", "existent_token2").Return(nil)
			},
		},
		{
			name:         "status ok",
			token:        "existent_token3",
			param:        `shortURL`,
			expectedCode: http.StatusNotFound,
			funcMock: func() {
				mockStorager.On("DeleteLinkByShortLink", mock_test.Anything, "shortURL", "existent_token3").Return(errorsapp.ErrNotFoundLink)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.funcMock != nil {
				tc.funcMock()
			}

			u, err := url.Parse("/links/short")
			if err != nil {
				fmt.Println(err)
				return
			}
			if tc.param != "" {
				q := u.Query()
				q.Add("url", tc.param)
				u.RawQuery = q.Encode()
			}
			req, err := http.NewRequest("DELETE", u.String(), nil)
			if err != nil {
				t.Fatal(err)
			}
			var tokenIDCTX models.CtxString = "Authorization"
			req = req.WithContext(context.WithValue(req.Context(), tokenIDCTX, tc.token))

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if rr.Code != tc.expectedCode {
				t.Errorf("expected status %d, got %d", tc.expectedCode, rr.Code)
			}
		})
	}
}
