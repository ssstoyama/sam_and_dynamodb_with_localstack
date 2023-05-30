package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"ssstoyama/local-serverless/client"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	c, err := client.NewDynamoDB(ctx)
	if err != nil {
		return nil, err
	}

	output, err := c.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("sample"),
	})
	if err != nil {
		return nil, err
	}
	var items []*struct {
		ID string `json:"id" dynamodbav:"id"`
	}
	err = attributevalue.UnmarshalListOfMaps(output.Items, &items)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(items)
	if err != nil {
		return nil, err
	}
	return &events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(data)}, nil
}
