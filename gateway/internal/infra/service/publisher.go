package service

import (
	"context"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/service"
)

func NewPublisher() service.Publisher {
	return &publisherImpl{}
}

type publisherImpl struct{}

func (p *publisherImpl) PushObject(ctx context.Context, key string) error {
	return nil
}
