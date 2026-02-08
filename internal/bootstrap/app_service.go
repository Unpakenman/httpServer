package bootstrap

import (
	"context"
	"fmt"
	"httpServer/internal/app/client/http"
	"httpServer/internal/app/client/pg"
	"httpServer/internal/app/config"
	"httpServer/internal/app/httpserver"
	httpmapper "httpServer/internal/app/httpserver/mapper"
	ihttpservice "httpServer/internal/app/internal_services/internal_http_service"
	"httpServer/internal/app/provider"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
	"log/slog"
	"os"
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
	validatorInstance := validator.New()
	dbConn, err := pg.New(config.Config.ClinicsDB)
	if err != nil {
		log.Error(err.Error())
	}
	exit := make(chan os.Signal, 1)
	goExampleDBProvider := provider.NewGoExampleDBProvider(dbConn)

	httpClient := http.NewHTTPClient(config.Config.HttpClient, log)
	httpMapperInstance := httpmapper.New()
	someService := ihttpservice.NewService(config.Config.SomeHttpService, httpClient)
	someUseCase := clinics.NewUseCase(goExampleDBProvider, someService)
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

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error(err.Error(), "failed to shutdown http server")
	}

	if err := dbConn.CloseConnections(); err != nil {
		log.Error(err.Error(), "failed to close database connection")
	}

}
