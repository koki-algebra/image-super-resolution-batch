package service

import (
	"context"
	"io"
)

type Storage interface {
	PutObject(ctx context.Context, key string, r io.Reader) error
}
