package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/caiquetorres/amparo/api"
)

func main() {
	server := api.NewServer()
	adapter := gorillamux.NewV2(server.Router)
	lambda.Start(adapter.ProxyWithContext)
}
