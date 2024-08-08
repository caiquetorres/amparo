package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/caiquetorres/amparo/cmd/api/middleware"
	"github.com/caiquetorres/amparo/cmd/api/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()

	router.Use(middleware.Logging)
	router.Use(middleware.Cors)

	routes.SetupNotFoundRoutes(router)
	routes.SetupPingRoutes(subRouter)
	routes.SetupImportantDatesRoutes(subRouter)

	adapter := gorillamux.NewV2(router)
	lambda.Start(adapter.ProxyWithContext)
}
