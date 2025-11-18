package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables from the
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Couldn't load env from the profile")
	}

	cfg := config{
		addr: ":8000",
		db: dbConfig{
			dsn: os.Getenv("DATABASE_URL"),
		},
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.db.dsn)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		panic(err)
	}
	defer conn.Close(ctx)

	log.Println("Connected to database successfuly!")

	api := application{
		config: cfg,
		db:     conn,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	var h = api.mount()
	if err := api.run(h); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
