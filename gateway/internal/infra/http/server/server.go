package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/database"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/service"
	"github.com/sourcegraph/conc/pool"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Connect to database
	sqlDB, err := database.Open(ctx, s.cfg)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	// Initialize AWS config
	awsCfg, err := service.InitAWSConfig(ctx, s.cfg)
	if err != nil {
		return err
	}

	// router
	router, err := newRouter(s.cfg, sqlDB, awsCfg)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Handler:           router,
		Addr:              fmt.Sprintf(":%d", s.cfg.ServerPort),
		WriteTimeout:      time.Second * 60,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		IdleTimeout:       time.Second * 120,
	}

	pool := pool.New().WithErrors().WithContext(ctx)
	pool.Go(func(ctx context.Context) error {
		slog.Info(fmt.Sprintf("start HTTP server port: %d", s.cfg.ServerPort))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	<-ctx.Done()
	slog.Info("stopping HTTP server...")
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	return pool.Wait()
}
