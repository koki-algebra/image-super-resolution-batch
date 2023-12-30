package service

import "context"

type Publisher interface {
	PushObject(ctx context.Context, key string) error
}
