package client

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// 環境変数から値を取得して aws.Config 構造体を生成する
//
// 環境変数:
// - AWS_REGION
// - AWS_ENDPOINT
func newConfig(ctx context.Context) (aws.Config, error) {
	awsRegion := ""
	if region := os.Getenv("AWS_REGION"); region != "" {
		awsRegion = region
	}
	awsEndpoint := ""
	if endpoint := os.Getenv("AWS_ENDPOINT"); endpoint != "" {
		awsEndpoint = endpoint
	}
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// デフォルト設定にフォールバックする
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	return config.LoadDefaultConfig(ctx,
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
}
