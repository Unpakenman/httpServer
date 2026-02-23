package client

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	client       *Client
	exchangeName string

	channel *amqp.Channel
}

func NewPublisher(client *Client, exchangeName string) (*Publisher, error) {
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}

	p := Publisher{
		client:       client,
		exchangeName: exchangeName,
	}

	err := p.connect()
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Publisher) connect() error {
	if p.channel != nil && !p.channel.IsClosed() {
		return nil
	}

	if p.client.Conn == nil || p.client.Conn.IsClosed() {
		return fmt.Errorf("no rmq connection")
	}

	ch, err := p.client.Conn.Channel()
	if err != nil {
		return err
	}
	p.channel = ch

	err = p.channel.ExchangeDeclarePassive(
		p.exchangeName,
		"",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Publisher) Publish(ctx context.Context, routingKey string, data amqp.Publishing) error {
	err := p.connect()
	if err != nil {
		return err
	}

	err = p.channel.PublishWithContext(
		ctx,
		p.exchangeName,
		routingKey,
		false,
		false,
		data,
	)
	if err != nil {
		return err
	}

	p.client.Logger.InfoCtx(ctx, fmt.Sprintf(
		"Sent message: %s to exchange : \"%s\", rouning key: \"%s\"",
		string(data.Body),
		p.exchangeName,
		routingKey,
	))

	return nil
}

func (p *Publisher) Close() error {
	if p.channel != nil {
		return p.channel.Close()
	}

	return nil
}
