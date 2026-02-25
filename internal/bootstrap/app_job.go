package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"httpServer/internal/app/client/http"
	"httpServer/internal/app/client/pg"
	"httpServer/internal/app/config"
	ihttpservice "httpServer/internal/app/internal_services/internal_http_service"
	"httpServer/internal/app/jobs"
	logger "httpServer/internal/app/log"
	"httpServer/internal/app/provider"
	"httpServer/internal/app/rabbitmq_service"
	rmqclient "httpServer/internal/app/rabbitmq_service/client"
	"httpServer/internal/app/usecase/clinics"
)

func RunJob(ctx context.Context, cfg *config.Values, logger logger.LogClient, jobName string) {
	dbConn, err := pg.New(cfg.ClinicsDB)
	if err != nil {
		logger.Fatal(err)
	}

	rmqClientStop := make(chan bool)
	rmqClient, err := NewRMQClient(rmqClientStop, cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	appointmentsPublisher, err := rmqclient.NewPublisher(rmqClient, cfg.AMQPServer.AppointmentsCommandsExchange)
	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		if err := dbConn.CloseConnections(); err != nil {
			logger.Error(fmt.Errorf("error closing connections: %w", err))
		}
	}()

	//mapperInstance := mapper.New()

	DBProvider := provider.NewGoExampleDBProvider(dbConn)
	appointmentRMQService := rabbitmq_service.NewRMQService(appointmentsPublisher)

	httpClient := http.NewHTTPClient(cfg.HttpClient, logger)

	someService := ihttpservice.NewService(cfg.SomeHttpService, httpClient)

	clinicUseCase := clinics.NewUseCase(DBProvider, logger, someService, appointmentRMQService, cfg)

	jobsFn := jobs.NewJobList(clinicUseCase)

	for _, j := range jobsFn {
		if j.GetName() == jobName {
			logger.Info(fmt.Sprintf("job.start %s", jobName))
			err := j.Run(ctx)
			if err != nil {
				logger.Fatal(errors.Join(fmt.Errorf("job %s failed", jobName), err))
			}
			logger.Info(fmt.Sprintf("job.finish %s", jobName))
			return
		}
	}
	logger.Fatal(fmt.Errorf("job %s not found", jobName))
}
