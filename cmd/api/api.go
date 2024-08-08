package api

import (
	"fmt"
	"net/http"

	"github.com/caiquetorres/amparo/cmd/api/middleware"
	"github.com/caiquetorres/amparo/cmd/api/routes"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api").Subrouter()

	router.Use(middleware.Logging)
	router.Use(middleware.Cors)

	routes.SetupNotFoundRoutes(router)
	routes.SetupPingRoutes(subRouter)
	routes.SetupImportantDatesRoutes(subRouter)

	return &Server{Router: router}
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", addr), s.Router)
}
