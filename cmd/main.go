package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/nnaxim/go-backend-study/internal/env"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db:   dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=study sslmode=disable"),
		},
	}

	// db
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		fmt.Printf("database has failed to start, error: %s", err)
		os.Exit(1)
	}

	defer conn.Close(ctx)

	fmt.Printf("connected to database: %s", cfg.db.dsn)

	api := application{
		config: cfg,
		db: conn,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		fmt.Printf("server has failed to start, error: %s", err)
		os.Exit(1)
	}

}