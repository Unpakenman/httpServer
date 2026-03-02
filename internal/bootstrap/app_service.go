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
	"httpServer/internal/app/client/redis"
	"httpServer/internal/app/config"
	eventshandlers "httpServer/internal/app/events/handlers"
	"httpServer/internal/app/grpcserver"
	"httpServer/internal/app/grpcserver/mapper"
	"httpServer/internal/app/httpserver"
	httpmapper "httpServer/internal/app/httpserver/mapper"
	ihttpservice "httpServer/internal/app/internal_services/internal_http_service"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/provider"
	cacheprovider "httpServer/internal/app/provider/cache"
	"httpServer/internal/app/rabbitmq_service"
	rabbitmq "httpServer/internal/app/rabbitmq_service/client"
	"httpServer/internal/app/usecase/clinics"
	"httpServer/internal/app/validator"
	"os"
	"os/signal"
	"syscall"
)

func RunService(ctx context.Context, cfg *config.Values, log logger.LogClient) {

	exit := make(chan os.Signal, 1)
	chiRouter := NewChiRouter()
	httpConfig := cfg.HttpServer

	httpServer, err := RunHTTPServer(chiRouter, log, httpConfig)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	mapperInstance := mapper.New()
	validatorInstance := validator.New()
	dbConn, err := pg.New(cfg.ClinicsDB)
	if err != nil {
		log.Error(err)
	}

	goExampleDBProvider := provider.NewGoExampleDBProvider(dbConn)

	rmqClientStop := make(chan bool)
	rmqClient, err := NewRMQClient(rmqClientStop, cfg, log)
	if err != nil {
		log.Fatal(err)
	}
	rmqSms, err := rabbitmq.NewPublisher(rmqClient, cfg.AMQPServer.SmsQueue)
	if err != nil {
		log.Fatal(err)
	}
	rmqService := rabbitmq_service.NewRMQService(rmqSms)

	cacheClient, err := redis.NewRedisClient(cfg.Redis.URL)
	if err != nil {
		log.Fatal(err)
	}
	cacheProvider := cacheprovider.NewRedisCache(cacheClient, cfg.Redis.Prefix)

	httpClient := http.NewHTTPClient(cfg.HttpClient, log)
	httpMapperInstance := httpmapper.New()

	someService := ihttpservice.NewService(cfg.SomeHttpService, httpClient)
	clinicsUseCaseInstance := clinics.NewUseCase(goExampleDBProvider, log, someService, rmqService, cfg, cacheProvider)

	eventHandlers := eventshandlers.NewHandlerList(log, clinicsUseCaseInstance)

	grpcPortListener, err := NewGRPCPortListener(cfg.GRPCServer)
	if err != nil {
		log.Error(err)
	}
	defer func() {
		err := grpcPortListener.Close()
		if err != nil {
			log.Error(err)
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
		log.Error(err)
	}

	pb.RegisterClinicsServer(grpcServer, clinicServerInstance)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthcheck)
	reflection.Register(grpcServer)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := grpcServer.Serve(grpcPortListener); err != nil {
			log.Error(err, "grpc serve failed")
		}
	}()

	stopEventListener := RunEventListner(
		rmqClient,
		*cfg,
		log,
		eventHandlers)

	_ = httpserver.NewHttpServer(
		log,
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
		log.InfoCtx(ctx, "ctx.Done: ", done)
	}
	grpcServer.GracefulStop()
	stopEventListener()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error(err, "failed to shutdown http server")
	}

	if err := dbConn.CloseConnections(); err != nil {
		log.Error(err, "failed to close database connection")
	}
	log.Info("Server Exited Properly")
}
