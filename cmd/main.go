package main

import (
	"github.com/caiquetorres/amparo/api"
	"github.com/caiquetorres/amparo/config"
)

func main() {
	server := api.NewServer()
	server.ListenAndServe(config.Envs.Port)
}
