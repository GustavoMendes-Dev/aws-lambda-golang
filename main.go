package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received request:", request)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, AWS Lambda with Go!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
