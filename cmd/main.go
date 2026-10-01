package main

import (
	"log/slog"
	"os"
)

func main() {
	config := config{
		addr: "localhost:8080",
		db:   dbConfig{},
	}

	api := application{
		config: config,
	}

	// Logger
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	if err := api.run(api.mount()); err != nil {
		slog.Error("服务启动失败", "error", err)
		os.Exit(1)
	}
}
