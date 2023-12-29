package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/infra/http/server"
)

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error("failed to terminated server", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.New()
	if err != nil {
		return err
	}

	srv := server.NewServer(cfg)

	return srv.Run(ctx)
}
