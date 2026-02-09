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

func RunService(ctx context.Context) {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)
	chiRouter := NewChiRouter()
	httpConfig := config.Config.HttpServer

	httpServer, err := RunHTTPServer(chiRouter, httpConfig)
	if err != nil {
		log.Error(err.Error())
	}
	mapperInstance := mapper.New()
	validatorInstance := validator.New()
	dbConn, err := pg.New(config.Config.ClinicsDB)
	if err != nil {
		log.Error(err.Error())
	}

	goExampleDBProvider := provider.NewGoExampleDBProvider(dbConn)
	httpClient := http.NewHTTPClient(config.Config.HttpClient, log)
	httpMapperInstance := httpmapper.New()
	someService := ihttpservice.NewService(config.Config.SomeHttpService, httpClient)
	someUseCase := clinics.NewUseCase(goExampleDBProvider, someService)
	grpcPortListener, err := NewGRPCPortListener(config.Config.GRPCServer)
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
		validatorInstance,
		mapperInstance,
		someUseCase)
	healthcheck := health.NewServer()
	grpcServer, err := NewGRPCServer(config.Config.GRPCServer, log)
	if err != nil {
		log.Error(err.Error())
	}
	pb.RegisterClinicsServer(grpcServer, clinicServerInstance)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthcheck)
	reflection.Register(grpcServer)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		err = grpcServer.Serve(grpcPortListener)
		if err != nil {
			log.Error(err.Error())
		}
	}()

	_ = httpserver.NewHttpServer(
		*log,
		chiRouter,
		config.Config.HttpServer,
		httpMapperInstance,
		validatorInstance,
		someUseCase,
	)
	log.Info("app is ready")
	select {
	case v := <-exit:
		log.Warn(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		log.Error("ctx.Done: %v", done)
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
