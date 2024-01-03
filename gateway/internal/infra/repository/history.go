package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
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
		ExcludeColumn("history_id", "timestamp").
		Returning("history_id, timestamp").
		Exec(ctx); err != nil {
		return err
	}

	return tx.Commit()
}

func (h *historyImpl) List(ctx context.Context, params repository.HistoryListParams) ([]*entity.History, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var limit, offset int
	if params.Limit > 0 {
		limit = params.Limit
	}
	if params.Offset > 0 {
		offset = params.Offset
	}

	var histories []*entity.History
	if err := tx.NewSelect().
		Model(&histories).
		Relation("IsrJob").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return histories, tx.Commit()
}

func (h *historyImpl) FindByJobID(ctx context.Context, jobID uuid.UUID) (*entity.History, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	history := new(entity.History)
	if err := tx.NewSelect().
		Model(history).
		Relation("IsrJob").
		Where("history.isr_job_id = ?", jobID).
		Order("timestamp DESC").
		Limit(1).
		Scan(ctx); err != nil {
		return nil, err
	}

	return history, tx.Commit()
}
