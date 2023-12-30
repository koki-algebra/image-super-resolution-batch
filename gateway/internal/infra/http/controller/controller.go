package controller

import (
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/oapi"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/usecase"
)

func New(img usecase.Image) oapi.ServerInterface {
	return &controllerImpl{
		image: img,
	}
}

type controllerImpl struct {
	image usecase.Image
}
