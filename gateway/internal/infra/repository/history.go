package repository

import (
	"context"
	"database/sql"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/entity"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/repository"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewHistory(sqlDB *sql.DB) repository.History {
	return &historyImpl{
		db: bun.NewDB(sqlDB, pgdialect.New()),
	}
}

type historyImpl struct {
	db *bun.DB
}

func (h *historyImpl) Create(ctx context.Context, history *entity.History) error {
	return nil
}
