package main

import (
	"context"
	"httpServer/internal/app/config"
	appLog "httpServer/internal/app/log"
	"httpServer/internal/bootstrap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	logger, err := appLog.New(*cfg)
	config.Config = cfg
	bootstrap.RunService(ctx, cfg, logger)
}
