package controller

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/oapi"
)

func renderMessage(w http.ResponseWriter, r *http.Request, code int, msg string) {
	render.Status(r, code)
	render.JSON(w, r, oapi.Message{Message: &msg})
}
