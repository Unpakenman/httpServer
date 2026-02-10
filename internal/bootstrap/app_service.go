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

func RunService(ctx context.Context, cfg *config.Values, logger slog.Logger) {

	chiRouter := NewChiRouter()
	httpConfig := cfg.HttpServer

	httpServer, err := RunHTTPServer(chiRouter, logger, httpConfig)
	if err != nil {
		logger.Error(err.Error())
	}

	mapperInstance := mapper.New()
	validatorInstance := validator.New()
	dbConn, err := pg.New(cfg.ClinicsDB)
	if err != nil {
		logger.Error(err.Error())
	}

	goExampleDBProvider := provider.NewGoExampleDBProvider(dbConn)

	httpClient := http.NewHTTPClient(cfg.HttpClient, &logger)
	httpMapperInstance := httpmapper.New()

	someService := ihttpservice.NewService(cfg.SomeHttpService, httpClient)
	clinicsUseCaseInstance := clinics.NewUseCase(goExampleDBProvider, logger, someService, cfg)

	grpcPortListener, err := NewGRPCPortListener(cfg.GRPCServer)
	if err != nil {
		logger.Error(err.Error())
	}
	defer func() {
		err := grpcPortListener.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	clinicServerInstance := grpcserver.NewClinicServer(
		&logger,
		validatorInstance,
		mapperInstance,
		clinicsUseCaseInstance)
	healthcheck := health.NewServer()
	grpcServer, err := NewGRPCServer(cfg.GRPCServer, &logger)
	if err != nil {
		logger.Error(err.Error())
	}
	pb.RegisterClinicsServer(grpcServer, clinicServerInstance)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthcheck)
	reflection.Register(grpcServer)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		err = grpcServer.Serve(grpcPortListener)
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	_ = httpserver.NewHttpServer(
		logger,
		chiRouter,
		cfg.HttpServer,
		httpMapperInstance,
		validatorInstance,
		clinicsUseCaseInstance,
	)
	fmt.Println("app service is running")
	select {
	case v := <-exit:
		logger.Warn(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		logger.Error("ctx.Done: %v", done)
	}
	grpcServer.GracefulStop()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error(err.Error(), "failed to shutdown http server")
	}

	if err := dbConn.CloseConnections(); err != nil {
		logger.Error(err.Error(), "failed to close database connection")
	}
	logger.Info("Server Exited Properly")
}
