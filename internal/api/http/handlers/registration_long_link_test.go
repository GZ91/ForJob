package handlers

import (
	"context"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/GZ91/linkreduct/internal/service"
	"github.com/GZ91/linkreduct/internal/service/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_handlers_AddLongLinkJSON(t *testing.T) {
	logger.Initializing("info")

	mockStorager := &mocks.Storeger{}
	configa := config.New(false, "", "", 100, 4, "", "existent_token")

	serviceNode := service.New(context.Background(), mockStorager, configa)

	h := &handlers{
		conf:        configa, // замените на вашу реальную структуру конфигурации
		nodeService: serviceNode,
	}

	r := chi.NewRouter()
	r.HandleFunc("/shortlink", h.AddLongLinkJSON)

	type TestCase struct {
		name         string
		token        string
		longLink     string
		expectedCode int
		body         string
		funcMock     func()
	}

	testCases := []TestCase{
		{
			name:     "new link",
			longLink: "/test_long.com",
			body:     `{ "longLink": "http://test_long.com"}`,
			funcMock: func() {
				mockStorager.On("FindLongURL", mock.Anything, "http://test_long.com", "token").Return("test_short.com", false, nil)
				mockStorager.On("AddURL", mock.Anything, "http://test_long.com").Return("test_short.com", nil)
			},
			expectedCode: http.StatusOK,
			token:        "token",
		},
		{
			name:     "old link",
			longLink: "/test_long.com",
			body:     `{ "longLink": "http://test_long.com"}`,
			funcMock: func() {
				mockStorager.On("FindLongURL", mock.Anything, "http://test_long.com", "token").Return("test_short.com", true, nil)
				mockStorager.On("AddURL", mock.Anything, "http://test_long.com").Return("test_short.com", nil)
			},
			expectedCode: http.StatusOK,
			token:        "token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.funcMock != nil {
				tc.funcMock()
			}

			req, err := http.NewRequest("GET", "/shortlink", strings.NewReader(tc.body))
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
