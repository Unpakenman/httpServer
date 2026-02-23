package bootstrap

import (
	"fmt"
	"httpServer/internal/app/config"
	"httpServer/internal/app/constants"
	"httpServer/internal/app/events"
	logger "httpServer/internal/app/log"
	rabbitmq "httpServer/internal/app/rabbitmq_service/client"
	amqplistener "httpServer/internal/app/rabbitmq_service/listener"
	"httpServer/internal/app/rabbitmq_service/otelamqp"
	"log"
	"sync"
)

func RunEventListner(
	client *rabbitmq.Client,
	cfg config.Values,
	logger logger.LogClient,
	handlersList []events.Handler,
) func() {
	consumerProvider, err := otelamqp.NewConsumerOtelProvider(
		otelamqp.OtelConfig{
			QueueName:    cfg.AMQPServer.EventsQueue,
			ProtocolName: cfg.AMQPServer.Protocol,
			ClientID:     cfg.App.Name,
			ServerName:   cfg.AMQPServer.Hostname,
			ServerPort:   int(cfg.AMQPServer.Port),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	router := events.NewRouter(
		handlersList,
		logger,
		consumerProvider,
	)

	var wg sync.WaitGroup

	wg.Add(1)
	listenerCfg := &amqplistener.ListenerConfig{
		QueueName:     cfg.AMQPServer.EventsQueue,
		PrefetchCount: cfg.AMQPServer.EventsPrefetchCount,
	}
	listener, err := amqplistener.NewListenerWithRoutingKeyFunc(client, listenerCfg, router.Process)
	if err != nil {
		logger.Fatal(fmt.Errorf(`%s: %w`, constants.RabbitMQLabel, err))
	}

	go func() {
		err := listener.Run()
		if err != nil {
			logger.Fatal(fmt.Errorf(`%s: %w`, constants.RabbitMQLabel, err))
		}
		wg.Done()
	}()
	return func() {
		listener.Close()
		wg.Wait()
	}
}
