package main

import (
	"github.com/caiquetorres/amparo/cmd/api"
	_ "github.com/caiquetorres/amparo/cmd/docs"
	"github.com/caiquetorres/amparo/config"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Amparo API
// @version 1.0
func main() {
	server := api.NewServer()
	server.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	server.ListenAndServe(config.Envs.Port)
}
