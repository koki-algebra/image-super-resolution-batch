package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/controller"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/oapi"
)

func newRouter(sqlDB *sql.DB) (http.Handler, error) {
	r := chi.NewRouter()

	swagger, err := oapi.GetSwagger()
	if err != nil {
		return nil, err
	}
	swagger.Servers = nil

	// logger
	logger := httplog.NewLogger("app", httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(logger))

	ctrl := controller.New()
	oapi.HandlerFromMux(ctrl, r)

	return r, nil
}
