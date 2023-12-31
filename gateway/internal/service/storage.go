package service

import (
	"context"
	"io"
)

type Storage interface {
	PutObject(ctx context.Context, bucket string, key string, data io.Reader) error
}
