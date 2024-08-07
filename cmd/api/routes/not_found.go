package routes

import (
	"net/http"

	handlers "github.com/caiquetorres/amparo/cmd/api/hadlers"
	"github.com/gorilla/mux"
)

func SetupNotFoundRoutes(router *mux.Router) {
	handler := handlers.NewNotFoundHandler()
	router.NotFoundHandler = http.HandlerFunc(handler.HandleNotFound)
}
