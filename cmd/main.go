package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caiquetorres/amparo/cmd/api"
	"github.com/caiquetorres/amparo/config"
)

func main() {
	server := api.NewServer(fmt.Sprintf(":%s", config.Envs.Port))
	go func() {
		if err := server.Run(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop
	log.Println("Server stopped")
}
