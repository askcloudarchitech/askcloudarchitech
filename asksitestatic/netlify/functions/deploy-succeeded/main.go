package main

import (
	"context"

	"github.com/askcloudarchitech/mediumautopost/pkg/mediumautopost"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	mediumautopost.Do("")

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Body:            "Success",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
