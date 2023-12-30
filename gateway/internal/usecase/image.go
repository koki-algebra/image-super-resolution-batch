package usecase

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/entity"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/repository"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/service"
)

type Image interface {
	Publish(ctx context.Context, r io.Reader, extension string) (*entity.History, error)
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

func (i *imageImpl) Publish(ctx context.Context, r io.Reader, extension string) (*entity.History, error) {
	jobID := uuid.NewString()
	history := entity.History{
		Status: entity.PENDING,
		IsrJob: entity.IsrJob{
			IsrJobID:                jobID,
			UploadImageKey:          fmt.Sprintf("%s/%s%s", i.cfg.UploadImagePrefix, jobID, extension),
			SuperResolutionImageKey: fmt.Sprintf("%s/%s%s", i.cfg.SuperResolutionImagePrefix, jobID, extension),
		},
	}

	// Uploaded image's object key
	imageKey := history.IsrJob.UploadImageKey

	// Create history
	if err := i.repo.Create(ctx, &history); err != nil {
		return nil, err
	}

	// Upload the uploaded image to object storage
	if err := i.storage.PutObject(ctx, imageKey, r); err != nil {
		return nil, err
	}

	// Push image key to messaging queue
	if err := i.publisher.PushObject(ctx, imageKey); err != nil {
		return nil, err
	}

	return nil, nil
}
