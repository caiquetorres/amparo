package routes

import (
	"log"

	handlers "github.com/caiquetorres/amparo/cmd/api/hadlers"
	"github.com/gorilla/mux"
)

func SetupPingRoutes(router *mux.Router) {
	handler := handlers.NewPingHandler()
	router.HandleFunc("/ping", handler.HandlePingGet).Methods("GET")
	log.Println("Mapped GET /ping")
}
