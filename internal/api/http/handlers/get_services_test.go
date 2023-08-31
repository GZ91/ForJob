package handlers

import (
	"context"
	"errors"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/GZ91/linkreduct/internal/service"
	"github.com/GZ91/linkreduct/internal/service/mocks"
	"github.com/go-chi/chi/v5"
	mock_test "github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_handlers_GetServices(t *testing.T) {
	logger.Initializing("info")

	mockStorager := &mocks.Storeger{}
	configa := config.New(false, "", "", 100, 4, "", "existent_token")

	serviceNode := service.New(context.Background(), mockStorager, configa)

	h := &handlers{
		conf:        configa, // замените на вашу реальную структуру конфигурации
		nodeService: serviceNode,
	}

	r := chi.NewRouter()
	r.HandleFunc("/services", h.GetServices)

	type TestCase struct {
		name         string
		token        string
		body         string
		expectedCode int
		funcMock     func()
	}

	testCases := []TestCase{
		{
			name:         "Unauthorized",
			token:        "other",
			body:         "",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "not json",
			token:        "existent_token",
			body:         "other text not json",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "error when accessing the get service function",
			token:        "existent_token",
			body:         `{"name": "nameService"}`,
			expectedCode: http.StatusInternalServerError,
			funcMock: func() {
				mockStorager.On("GetServices", mock_test.Anything, "nameService").Return(nil, errors.New("other"))
			},
		},
		{
			name:         "status ok",
			token:        "existent_token",
			body:         `{"name": "nameService3"}`,
			expectedCode: http.StatusOK,
			funcMock: func() {
				data := make(map[string]string)
				data["nameService"] = "qwer-tyuiop974u-h89u"
				mockStorager.On("GetServices", mock_test.Anything, "nameService3").Return(data, nil)
			},
		},
		{
			name:         "status ok2",
			token:        "existent_token",
			body:         "",
			expectedCode: http.StatusOK,
			funcMock: func() {
				data := make(map[string]string)
				data["nameService"] = "qwer-tyuiop974u-h89u"
				mockStorager.On("GetServices", mock_test.Anything, "").Return(data, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("POST", "/services", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			if tc.funcMock != nil {
				tc.funcMock()
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
