package service

import (
	"context"
	"io"
)

type Storage interface {
	GetObject(ctx context.Context, bucket string, key string) (io.ReadCloser, error)
	PutObject(ctx context.Context, bucket string, key string, data io.Reader) error
}
