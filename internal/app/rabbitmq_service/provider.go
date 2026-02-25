package rabbitmq_service

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"httpServer/internal/app/rabbitmq_service/client"
	"time"
)

type rabbitmqService struct {
	smsPublisher *client.Publisher
}

type RMQService interface {
	SendAppointmentsMessage(ctx context.Context, req []AppointmentsListMessage) error
}

func NewRMQService(
	smsPublisher *client.Publisher,
) RMQService {
	return &rabbitmqService{
		smsPublisher: smsPublisher,
	}
}

func (s *rabbitmqService) PublishMessage(ctx context.Context, publisher *client.Publisher, message []byte) error {
	err := publisher.Publish(ctx, "", amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		Timestamp:    time.Now(),
		Body:         message,
	})

	return err
}
