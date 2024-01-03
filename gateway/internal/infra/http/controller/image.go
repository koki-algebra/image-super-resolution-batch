package controller

import (
	"fmt"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/go-chi/render"
	openapi_types "github.com/oapi-codegen/runtime/types"
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

func (ctrl *controllerImpl) DownloadImage(w http.ResponseWriter, r *http.Request, jobID openapi_types.UUID) {
	ctx := r.Context()

	img, ext, err := ctrl.image.Download(ctx, jobID)
	if err != nil {
		slog.Error("failed to download image", "error", err)
		renderMessage(w, r, http.StatusInternalServerError, "internal error")
		return
	}
	defer img.Close()

	// Set the content type header based on the image file extension
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)

	// Set the Content-Disposition header to suggest a filename
	filename := fmt.Sprintf("%s%s", jobID.String(), ext)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))

	// Copy the image data to the response writer
	if _, err := io.Copy(w, img); err != nil {
		slog.Error("failed to download image", "error", err)
		renderMessage(w, r, http.StatusInternalServerError, "internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
}
