package service

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/service"
)

func NewStorage(cfg aws.Config) service.Storage {
	client := s3.NewFromConfig(cfg)
	return &storageImpl{
		client: client,
	}
}

type storageImpl struct {
	client *s3.Client
}

func (s *storageImpl) GetObject(ctx context.Context, bucket string, key string) (io.ReadCloser, error) {
	res, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func (s *storageImpl) PutObject(ctx context.Context, bucket string, key string, data io.Reader) error {
	if _, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   data,
	}); err != nil {
		return err
	}

	return nil
}
