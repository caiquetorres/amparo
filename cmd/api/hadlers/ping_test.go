package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/caiquetorres/amparo/cmd/api/hadlers"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/api/ping", nil)
		w := httptest.NewRecorder()

		handler := handlers.NewPingHandler()
		handler.HandlePingGet(w, r)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "pong", string(body))
	})
}
