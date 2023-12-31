package controller

import (
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/oapi"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/usecase"
)

func New(img usecase.Image, history usecase.History) oapi.ServerInterface {
	return &controllerImpl{
		image:   img,
		history: history,
	}
}

type controllerImpl struct {
	image   usecase.Image
	history usecase.History
}
