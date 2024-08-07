package handlers

import (
	"net/http"

	"github.com/caiquetorres/amparo/internal"
)

type NotFoundHandler struct{}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (h *NotFoundHandler) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	internal.WriteJSON(w, http.StatusNotFound, internal.NewApiError("Not found"))
}
