package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Joskmo/ecom-golang-study.git/internal/env"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Context
	ctx := context.Background()

	// Loading config
	if err := godotenv.Load(); err != nil {
		slog.Warn("the .env file wasn't read -> using default data", "error", err)
	}
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString(
				"GOOSE_DBSTRING",
				"host=localhost user=postgres password=postgres dbname=ecom sslmode=disabled",
			),
		},
	}

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("failed to connect to the database", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	slog.Info("connected to database", "dsn", cfg.db.dsn)

	// Application
	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to starts", "error", err)
		os.Exit(1)
	}
}
