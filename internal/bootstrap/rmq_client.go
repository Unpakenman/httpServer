package bootstrap

import (
	"httpServer/internal/app/config"
	logger "httpServer/internal/app/log"
	rabbitmq "httpServer/internal/app/rabbitmq_service/client"
)

func NewRMQClient(
	clientStop chan bool,
	cfg *config.Values,
	logger logger.LogClient,
) (*rabbitmq.Client, error) {
	clientConfig := rabbitmq.AMQPConfig{
		User:             cfg.AMQPServer.User,
		Password:         cfg.AMQPServer.Password,
		Hostname:         cfg.AMQPServer.Hostname,
		Protocol:         cfg.AMQPServer.Protocol,
		VHost:            cfg.AMQPServer.VHost,
		Port:             cfg.AMQPServer.Port,
		ReconnectTimeout: cfg.AMQPServer.ReconnectTimeout,
	}

	return rabbitmq.NewClient(&clientConfig, logger, clientStop)
}
