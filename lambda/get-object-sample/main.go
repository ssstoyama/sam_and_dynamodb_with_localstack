package main

import (
	"bytes"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"ssstoyama/local-serverless/client"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	c, err := client.NewS3(ctx)
	if err != nil {
		return nil, err
	}

	output, err := c.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String("sample-bucket"),
		Key:    aws.String("sample.txt"),
	})
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(output.Body)
	if err != nil {
		return nil, err
	}
	return &events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: buf.String()}, nil
}
