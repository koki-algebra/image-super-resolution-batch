package usecase

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/entity"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/repository"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/service"
)

type Image interface {
	Publish(ctx context.Context, r io.Reader, extension string) (*entity.History, error)
	Download(ctx context.Context, jobID uuid.UUID) (r io.ReadCloser, ext string, err error)
}

func NewImage(cfg *config.Config, publisher service.Publisher, storage service.Storage, repo repository.History) Image {
	return &imageImpl{
		cfg:       cfg,
		publisher: publisher,
		storage:   storage,
		repo:      repo,
	}
}

type imageImpl struct {
	cfg       *config.Config
	publisher service.Publisher
	storage   service.Storage
	repo      repository.History
}

func (img *imageImpl) Publish(ctx context.Context, r io.Reader, extension string) (*entity.History, error) {
	jobID := uuid.NewString()
	history := entity.History{
		Status:   entity.PENDING,
		IsrJobID: jobID,
		IsrJob: entity.IsrJob{
			IsrJobID:                jobID,
			UploadImageKey:          fmt.Sprintf("%s/%s%s", img.cfg.StorageUploadImagePrefix, jobID, extension),
			SuperResolutionImageKey: fmt.Sprintf("%s/%s%s", img.cfg.StorageSuperResolutionImagePrefix, jobID, extension),
		},
	}

	// Uploaded image's object key
	imageKey := history.IsrJob.UploadImageKey

	// Create history
	if err := img.repo.Create(ctx, &history); err != nil {
		return nil, err
	}

	// Upload the uploaded image to object storage
	if err := img.storage.PutObject(ctx, img.cfg.StorageBucket, imageKey, r); err != nil {
		return nil, err
	}

	// Push image key to messaging queue
	if err := img.publisher.Publish(ctx, "text/plain", []byte(jobID)); err != nil {
		return nil, err
	}

	return &history, nil
}

func (img *imageImpl) Download(ctx context.Context, jobID uuid.UUID) (io.ReadCloser, string, error) {
	history, err := img.repo.FindByJobID(ctx, jobID)
	if err != nil {
		return nil, "", err
	}

	body, err := img.storage.GetObject(ctx, img.cfg.StorageBucket, history.IsrJob.SuperResolutionImageKey)
	if err != nil {
		return nil, "", err
	}

	ext := filepath.Ext(history.IsrJob.SuperResolutionImageKey)

	return body, ext, nil
}
