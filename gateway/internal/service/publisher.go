package service

import "context"

type Publisher interface {
	Publish(ctx context.Context, contentType string, body []byte) error
}
