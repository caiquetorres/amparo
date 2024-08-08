package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/caiquetorres/amparo/cmd/api"
	_ "github.com/caiquetorres/amparo/cmd/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	server := api.NewServer()
	server.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	adapter := gorillamux.NewV2(server.Router)
	lambda.Start(adapter.ProxyWithContext)
}
