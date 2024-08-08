package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiquetorres/amparo/cmd/api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	t.Run("Route not found", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/not-found", nil)
		w := httptest.NewRecorder()

		handler := handlers.NewNotFoundHandler()
		handler.HandleNotFound(w, r)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.JSONEq(t, `{"message": "Not found"}`, string(body))
	})
}
