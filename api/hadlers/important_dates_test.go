package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/caiquetorres/amparo/api/dtos"
	handlers "github.com/caiquetorres/amparo/api/hadlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleImportantDatesPost(t *testing.T) {
	handler := handlers.NewImportantDatesHandler()

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Missing date_of_death property",
			requestBody:    struct{}{},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Missing date_of_death property"}`,
		},
		{
			name:           "Invalid JSON body",
			requestBody:    "invalid-json",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Invalid JSON body"}`,
		},
		{
			name: "Invalid date format",
			requestBody: dtos.GetImportantDates{
				DateOfDeath: "invalid-date",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Invalid date_of_death property"}`,
		},
		{
			name: "Date in the future",
			requestBody: dtos.GetImportantDates{
				DateOfDeath: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"message":"The date_of_death cannot be in the future"}`,
		},
		{
			name: "Valid date param",
			requestBody: dtos.GetImportantDates{
				DateOfDeath: "2024-02-15",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"name":"schedule_mass","date":"2024-02-18"},{"name":"register_deach","date":"2024-03-01"},{"name":"pension_request","date":"2024-05-15"},{"name":"insurange_claim","date":"2025-02-14"}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bodyBytes, err := json.Marshal(tt.requestBody)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/important-dates", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			handler.HandleImportantDatesPost(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.JSONEq(t, tt.expectedBody, string(body))
		})
	}
}
