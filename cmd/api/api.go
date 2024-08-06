package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/caiquetorres/amparo/internal"
	"github.com/caiquetorres/amparo/service"
	"github.com/gorilla/mux"
)

type Server struct {
	addr string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	apiSubRouter := router.PathPrefix("/api").Subrouter()

	service := service.NewHandler()

	router.Use(loggingMiddleware)
	service.RegisterRoutes(apiSubRouter)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusNotFound, internal.ApiError{Message: "Not found"})
	})

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}
