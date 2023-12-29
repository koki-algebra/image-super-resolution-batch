package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

func newRouter(sqlDB *sql.DB) (http.Handler, error) {
	r := chi.NewRouter()

	// logger
	logger := httplog.NewLogger("app", httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(logger))

	return r, nil
}
