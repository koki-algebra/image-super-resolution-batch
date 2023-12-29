package controller

import "github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/oapi"

func New() oapi.ServerInterface {
	return &controllerImpl{}
}

type controllerImpl struct{}
