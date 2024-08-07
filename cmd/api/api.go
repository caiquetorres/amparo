package api

import (
	"log"
	"net/http"

	"github.com/caiquetorres/amparo/cmd/api/middleware"
	"github.com/caiquetorres/amparo/cmd/api/routes"
	"github.com/gorilla/mux"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	apiSubRouter := router.PathPrefix("/api").Subrouter()

	router.Use(middleware.Logging)
	router.Use(middleware.Cors)

	routes.SetupNotFoundRoutes(router)
	routes.SetupPingRoutes(apiSubRouter)
	routes.SetupImportantDatesRoutes(apiSubRouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
