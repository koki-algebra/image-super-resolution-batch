package service

import (
	"context"
	"io"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/service"
)

func NewStorage() service.Storage {
	return &storageImpl{}
}

type storageImpl struct{}

func (s *storageImpl) PutObject(ctx context.Context, key string, r io.Reader) error {
	return nil
}
