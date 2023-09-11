package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/logger"
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

func Test_handlers_GetLinks(t *testing.T) {
	logger.Initializing("info")

	mockStorager := &mocks.Storeger{}
	configa := config.New(false, "", "", 100, 4, "", "existent_token")

	serviceNode := service.New(context.Background(), mockStorager, configa)

	h := &handlers{
		conf:        configa, // замените на вашу реальную структуру конфигурации
		nodeService: serviceNode,
	}

	r := chi.NewRouter()
	r.HandleFunc("/links", h.GetLinks)

	type TestCase struct {
		name         string
		token        string
		param        string
		expectedCode int
		funcMock     func()
	}

	testCases := []TestCase{
		{
			name:         "error when accessing the get links function",
			token:        "existent_token",
			param:        "",
			expectedCode: http.StatusInternalServerError,
			funcMock: func() {
				mockStorager.On("GetLinks", mock_test.Anything, "existent_token").Return(nil, errors.New("other"))
			},
		},
		{
			name:         "StatusNoContent",
			token:        "existent_token2",
			param:        `nameService`,
			expectedCode: http.StatusNoContent,
			funcMock: func() {
				mockStorager.On("GetLinks", mock_test.Anything, "existent_token2").Return(nil, nil)
			},
		},
		{
			name:         "status ok",
			token:        "existent_token3",
			param:        `nameService3`,
			expectedCode: http.StatusOK,
			funcMock: func() {
				data := make(map[string]string)
				data["shortlink"] = "longlink"
				mockStorager.On("GetLinks", mock_test.Anything, "existent_token3").Return(data, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.funcMock != nil {
				tc.funcMock()
			}

			u, err := url.Parse("/links")
			if err != nil {
				fmt.Println(err)
				return
			}
			if tc.param != "" {
				q := u.Query()
				q.Add("name", tc.param)
				u.RawQuery = q.Encode()
			}
			req, err := http.NewRequest("GET", u.String(), nil)
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
