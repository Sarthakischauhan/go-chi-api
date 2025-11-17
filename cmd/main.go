package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8000",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	var h = api.mount()
	if err := api.run(h); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
