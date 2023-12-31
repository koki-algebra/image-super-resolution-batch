package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
)

func InitAWSConfig(ctx context.Context, cfg *config.Config) (aws.Config, error) {
	// AWS config
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && len(cfg.S3Endpoint) > 0 {
			return aws.Endpoint{
				URL:               cfg.S3Endpoint,
				HostnameImmutable: true,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	awsCfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithEndpointResolverWithOptions(resolver))
	if err != nil {
		return aws.Config{}, err
	}

	return awsCfg, nil
}
