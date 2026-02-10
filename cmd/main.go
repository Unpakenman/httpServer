package main

import (
	"context"
	"httpServer/internal/app/config"
	"httpServer/internal/bootstrap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	config.Config = cfg
	bootstrap.RunService(ctx, cfg)
}
