package main

import (
	"context"
	"httpServer/internal/app/config"
	"httpServer/internal/bootstrap"
	"log/slog"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var logger slog.Logger
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	bootstrap.RunService(ctx, cfg, logger)
}
