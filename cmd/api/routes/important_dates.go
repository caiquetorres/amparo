package routes

import (
	"log"

	handlers "github.com/caiquetorres/amparo/cmd/api/handlers"
	"github.com/gorilla/mux"
)

func SetupImportantDatesRoutes(router *mux.Router) {
	handler := handlers.NewImportantDatesHandler()
	router.HandleFunc("/important-dates", handler.HandleImportantDatesPost).Methods("POST")
	log.Println("Mapped POST /important-dates")
}
