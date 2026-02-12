package bootstrap

import (
	"context"
	"fmt"
	pb "github.com/Unpakenman/protos/gen/go/sso"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"httpServer/internal/app/client/http"
	"httpServer/internal/app/client/pg"
	"httpServer/internal/app/config"
	"httpServer/internal/app/grpcserver"
	"httpServer/internal/app/grpcserver/mapper"
	"httpServer/internal/app/httpserver"
	httpmapper "httpServer/internal/app/httpserver/mapper"
	ihttpservice "httpServer/internal/app/internal_services/internal_http_service"
	"httpServer/internal/app/provider"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func RunService(ctx context.Context, cfg *config.Values) {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)
	exit := make(chan os.Signal, 1)
	chiRouter := NewChiRouter()
	httpConfig := cfg.HttpServer

	httpServer, err := RunHTTPServer(chiRouter, *log, httpConfig)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	mapperInstance := mapper.New()
	validatorInstance := validator.New()
	dbConn, err := pg.New(cfg.ClinicsDB)
	if err != nil {
		log.Error(err.Error())
	}

	goExampleDBProvider := provider.NewGoExampleDBProvider(dbConn)

	httpClient := http.NewHTTPClient(cfg.HttpClient, log)
	httpMapperInstance := httpmapper.New()

	someService := ihttpservice.NewService(cfg.SomeHttpService, httpClient)
	clinicsUseCaseInstance := clinics.NewUseCase(goExampleDBProvider, *log, someService, cfg)

	grpcPortListener, err := NewGRPCPortListener(cfg.GRPCServer)
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		err := grpcPortListener.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	clinicServerInstance := grpcserver.NewClinicServer(
		log,
		validatorInstance,
		mapperInstance,
		clinicsUseCaseInstance)
	healthcheck := health.NewServer()
	grpcServer, err := NewGRPCServer(cfg.GRPCServer, log)
	if err != nil {
		log.Error(err.Error())
	}
	pb.RegisterClinicsServer(grpcServer, clinicServerInstance)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthcheck)
	reflection.Register(grpcServer)

	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := grpcServer.Serve(grpcPortListener); err != nil {
			log.Error("grpc serve failed", "err", err)
		}
	}()

	_ = httpserver.NewHttpServer(
		*log,
		chiRouter,
		cfg.HttpServer,
		httpMapperInstance,
		validatorInstance,
		clinicsUseCaseInstance,
	)
	log.Info("app service started")
	select {
	case v := <-exit:
		log.Warn(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		log.InfoContext(ctx, "ctx.Done: ", done)
	}
	grpcServer.GracefulStop()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error(err.Error(), "failed to shutdown http server")
	}

	if err := dbConn.CloseConnections(); err != nil {
		log.Error(err.Error(), "failed to close database connection")
	}
	log.Info("Server Exited Properly")
}
