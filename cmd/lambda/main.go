package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/caiquetorres/amparo/cmd/api"
	_ "github.com/caiquetorres/amparo/cmd/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Amparo API
// @version 1.0
func main() {
	server := api.NewServer()
	server.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	adapter := gorillamux.NewV2(server.Router)
	lambda.Start(adapter.ProxyWithContext)
}
