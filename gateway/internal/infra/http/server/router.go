package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/controller"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/oapi"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/repository"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/service"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/usecase"
)

func newRouter(sqlDB *sql.DB, cfg *config.Config) (http.Handler, error) {
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

	// services
	publisher := service.NewPublisher()
	storage := service.NewStorage()

	// repositories
	historyRepo := repository.NewHistory(sqlDB)

	// usecases
	img := usecase.NewImage(cfg, publisher, storage, historyRepo)

	ctrl := controller.New(img)
	oapi.HandlerFromMux(ctrl, r)

	return r, nil
}
