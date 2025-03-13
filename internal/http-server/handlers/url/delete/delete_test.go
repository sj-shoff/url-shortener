package delete_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"url-shortener/internal/http-server/handlers/url/delete"
	"url-shortener/internal/http-server/handlers/url/delete/mocks"
	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/handlers/slogdiscard"
)

func TestDeleteHandler(t *testing.T) {
	cases := []struct {
		name         string
		alias        string
		respError    string
		mockError    error
		expectedCode int
	}{
		{
			name:  "Success",
			alias: "test_alias",
		},
		{
			name:         "Empty alias",
			alias:        "",
			respError:    "empty alias",
			expectedCode: http.StatusOK,
		},
		{
			name:         "DeleteURL Error",
			alias:        "test_alias",
			respError:    "internal error",
			mockError:    errors.New("unexpected error"),
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			urlDeleterMock := mocks.NewURLDeleter(t)

			// Устанавливаем поведение мока в зависимости от наличия ошибки
			if tc.mockError != nil {
				urlDeleterMock.On("DeleteURL", tc.alias).
					Return(tc.mockError).
					Once()
			} else if tc.respError == "" {
				urlDeleterMock.On("DeleteURL", tc.alias).
					Return(tc.mockError).
					Once()
			}

			handler := delete.New(slogdiscard.NewDiscardLogger(), urlDeleterMock)

			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/url/%s", tc.alias), nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("alias", tc.alias)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if tc.expectedCode != 0 {
				require.Equal(t, tc.expectedCode, rr.Code)
			} else {
				require.Equal(t, http.StatusOK, rr.Code)
			}

			body := rr.Body.String()

			var response resp.Response

			err := json.Unmarshal([]byte(body), &response)
			require.NoError(t, err)

			require.Equal(t, tc.respError, response.Error)
			urlDeleterMock.AssertExpectations(t)

		})
	}
}
