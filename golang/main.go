package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	fmt.Println("Hello World!")

	response = events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World!",
	}

	return response, err
}

func main() {
	lambda.Start(handler)
}
