package routes

import (
	"log"
	"net/http"

	handlers "github.com/caiquetorres/amparo/api/hadlers"
	"github.com/gorilla/mux"
)

func SetupNotFoundRoutes(router *mux.Router) {
	handler := handlers.NewNotFoundHandler()
	router.NotFoundHandler = http.HandlerFunc(handler.HandleNotFound)
	log.Println("Mapped Not Found")
}
