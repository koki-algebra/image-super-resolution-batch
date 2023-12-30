package repository

import (
	"context"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/entity"
)

type History interface {
	Create(ctx context.Context, history *entity.History) error
}
