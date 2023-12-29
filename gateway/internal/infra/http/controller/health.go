package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

func (ctrl *controllerImpl) HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"message": "OK",
	}
	render.JSON(w, r, res)
}
