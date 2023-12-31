package controller

import (
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/go-chi/render"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 * 100 // 100 MB
)

func (ctrl *controllerImpl) UploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		renderMessage(w, r, http.StatusBadRequest, "failed to parse form")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		renderMessage(w, r, http.StatusBadRequest, "failed to get file")
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)

	history, err := ctrl.image.Publish(ctx, file, ext)
	if err != nil {
		slog.Error("failed to publish", "error", err)
		renderMessage(w, r, http.StatusInternalServerError, "internal error")
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, convertHistory(history))
}
