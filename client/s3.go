package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewS3 S3 Client 生成
func NewS3(ctx context.Context) (*s3.Client, error) {
	cfg, err := newConfig(ctx)
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	}), nil
}
