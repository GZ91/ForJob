package handlers

import (
	"context"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/service"
	"github.com/GZ91/linkreduct/internal/service/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handlers_GetLongURL(t *testing.T) {
	logger.Initializing("info")

	mockStorager := &mocks.Storeger{}
	configa := config.New(false, "", "", 100, 4, "", "existent_token")

	serviceNode := service.New(context.Background(), mockStorager, configa)

	h := &handlers{
		conf:        configa, // замените на вашу реальную структуру конфигурации
		nodeService: serviceNode,
	}

	r := chi.NewRouter()
	r.HandleFunc("/{id}", h.GetLongURL)

	type TestCase struct {
		name         string
		token        string
		link         string
		expectedCode int
		funcMock     func()
	}
	testCases := []TestCase{
		{
			name: "redirect test",
			link: "/test_short.com",
			funcMock: func() {
				mockStorager.On("GetURL", mock.Anything, "test_short.com").Return("test_long.com", true, nil)
			},
			expectedCode: http.StatusTemporaryRedirect,
		},
		{
			name: "not found test",
			link: "/test_shor2.com",
			funcMock: func() {
				mockStorager.On("GetURL", mock.Anything, "test_shor2.com").Return("", false, nil)
			},
			expectedCode: http.StatusNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.funcMock != nil {
				tc.funcMock()
			}

			req, err := http.NewRequest("GET", tc.link, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			if rr.Code != tc.expectedCode {
				t.Errorf("expected status %d, got %d", tc.expectedCode, rr.Code)
			}

		})
	}
}
