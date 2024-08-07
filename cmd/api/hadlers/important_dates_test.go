package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	handlers "github.com/caiquetorres/amparo/cmd/api/hadlers"
	"github.com/stretchr/testify/assert"
)

func TestHandleImportantDatesPost(t *testing.T) {
	t.Run("Valid date param", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/important-dates", nil)
		q := r.URL.Query()
		q.Add("date", "2024-02-28")
		r.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()

		handler := handlers.NewImportantDatesHandler()
		handler.HandleImportantDatesPost(w, r)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.JSONEq(t, `{"schedule_mass": "2024-03-02","register_death": "2024-03-14","pension_request": "2024-05-28","insurance_claim": "2025-02-27"}`, string(body))
	})

	t.Run("Missing date param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/important-dates", nil)
		q := req.URL.Query()
		q.Add("date", "")
		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()

		handler := handlers.NewImportantDatesHandler()
		handler.HandleImportantDatesPost(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.JSONEq(t, `{"message":"Missing date"}`, string(body))
	})

	t.Run("Invalid date param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/important-dates", nil)
		q := req.URL.Query()
		q.Add("date", "invalid-date")
		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()

		handler := handlers.NewImportantDatesHandler()
		handler.HandleImportantDatesPost(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.JSONEq(t, `{"message":"Invalid date"}`, string(body))
	})

	t.Run("Date in the future", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/important-dates", nil)
		q := req.URL.Query()
		futureDate := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
		q.Add("date", futureDate)
		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()

		handler := handlers.NewImportantDatesHandler()
		handler.HandleImportantDatesPost(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.JSONEq(t, `{"message":"The date cannot be in the future"}`, string(body))
	})
}
