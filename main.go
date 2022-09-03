package main

import (
	"context"

	// "fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// func runRouter() {
// 	router := registerRoutes()

// 	router.Run(":8080")
// }

// https://github.com/awslabs/aws-lambda-go-api-proxy#gin
func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ginLambda := registerRoutes()

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(LambdaHandler)
	// fmt.Println(runRouter())
	// runRouter()
}
