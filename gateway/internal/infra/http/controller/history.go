package controller

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/entity"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/oapi"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/repository"
)

func (ctrl *controllerImpl) ListHistories(w http.ResponseWriter, r *http.Request, params oapi.ListHistoriesParams) {
	ctx := r.Context()

	var (
		limit  = 50
		offset = 0
		latest bool
	)
	if params.Limit != nil && *params.Limit > 0 && *params.Limit < limit {
		limit = *params.Limit
	}
	if params.Offset != nil && *params.Offset > 0 {
		offset = *params.Offset
	}
	if params.Latest != nil {
		latest = *params.Latest
	}

	p := repository.HistoryListParams{
		Limit:  limit + 1,
		Offset: offset,
		Latest: latest,
	}

	histories, err := ctrl.history.List(ctx, p)
	if err != nil {
		slog.Error("failed to get histories", "error", err)
		renderMessage(w, r, http.StatusInternalServerError, "internal error")
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, convertHistories(histories, limit))
}

func convertHistory(history *entity.History) *oapi.History {
	if history == nil {
		return nil
	}

	status := history.Status.String()
	timestamp := history.Timestamp.String()
	jobID := uuid.MustParse(history.IsrJobID)

	return &oapi.History{
		HistoryID: &history.HistoryID,
		IsrJobID:  &jobID,
		Status:    &status,
		Timestamp: &timestamp,
	}
}

func convertHistories(histories []*entity.History, limit int) *oapi.ListHistoriesResponse {
	if histories == nil {
		return nil
	}

	hasNext := false
	if len(histories) > limit {
		hasNext = true
		histories = histories[:len(histories)-1]
	}

	resHistories := []oapi.History{}
	for i := range histories {
		conv := convertHistory(histories[i])
		resHistories = append(resHistories, *conv)
	}

	return &oapi.ListHistoriesResponse{
		HasNext:   &hasNext,
		Histories: &resHistories,
	}
}
