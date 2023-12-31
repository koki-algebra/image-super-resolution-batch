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
		db: bun.NewDB(sqlDB, pgdialect.New(), bun.WithDiscardUnknownColumns()),
	}
}

type historyImpl struct {
	db *bun.DB
}

func (h *historyImpl) Create(ctx context.Context, history *entity.History) error {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.NewInsert().Model(&history.IsrJob).Exec(ctx); err != nil {
		return err
	}

	if _, err := tx.NewInsert().
		Model(history).
		ExcludeColumn("history_id", "timestamp", "isr_job").
		Returning("history_id, timestamp").
		Exec(ctx); err != nil {
		return err
	}

	return tx.Commit()
}
