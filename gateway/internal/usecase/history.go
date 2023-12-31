package usecase

import (
	"context"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/entity"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/repository"
)

type History interface {
	List(ctx context.Context, params repository.HistoryListParams) ([]*entity.History, error)
}

func NewHistory(cfg *config.Config, repo repository.History) History {
	return &historyImpl{
		cfg:  cfg,
		repo: repo,
	}
}

type historyImpl struct {
	cfg  *config.Config
	repo repository.History
}

func (h *historyImpl) List(ctx context.Context, params repository.HistoryListParams) ([]*entity.History, error) {
	histories, err := h.repo.List(ctx, params)
	if err != nil {
		return nil, err
	}

	return histories, nil
}
