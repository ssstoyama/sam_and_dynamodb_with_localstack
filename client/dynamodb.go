package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// NewDynamoDB DynamoDB Client 生成
func NewDynamoDB(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := newConfig(ctx)
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}
