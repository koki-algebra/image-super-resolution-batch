package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/koki-algebra/image-super-resolution-batch/gateway/internal/config"
	_ "github.com/lib/pq"
)

func Open(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	// connect to database
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Minute * 15)
	db.SetConnMaxLifetime(time.Hour * 12)
	db.SetMaxIdleConns(15)
	db.SetMaxOpenConns(20)

	// verify connection
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
