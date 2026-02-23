package listener

import (
	"fmt"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	amqpclient "httpServer/internal/app/rabbitmq_service/client"
)

type Listener struct {
	client *amqpclient.Client
	cfg    *ListenerConfig

	processor Processor

	processWithRoutingKeyFunc    func(routingKey string, data []byte) error
	hasProcessWithRoutingKeyFunc bool

	closeError        chan *amqp.Error
	channel           *amqp.Channel
	msgs              <-chan amqp.Delivery
	reconnectionError chan error
	quit              chan struct{}
}

type Processor interface {
	Process(data []byte) error
}

// NewListener create Listener instance
func NewListener(client *amqpclient.Client, cfg *ListenerConfig, processor Processor) (*Listener, error) {
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if cfg == nil {
		return nil, fmt.Errorf("invalid queue config")
	}

	return &Listener{
		client:    client,
		cfg:       cfg,
		processor: processor,
		quit:      make(chan struct{}),
	}, nil
}

func NewListenerWithRoutingKeyFunc(client *amqpclient.Client, cfg *ListenerConfig, processWithRoutingKeyFunc func(routingKey string, data []byte) error) (*Listener, error) {
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if cfg == nil {
		return nil, fmt.Errorf("invalid queue config")
	}

	return &Listener{
		client:                       client,
		cfg:                          cfg,
		processWithRoutingKeyFunc:    processWithRoutingKeyFunc,
		hasProcessWithRoutingKeyFunc: true,
		quit:                         make(chan struct{}),
	}, nil
}

func (l *Listener) connect() error {
	if l.channel != nil && !l.channel.IsClosed() {
		return nil
	}

	if l.client.Conn == nil || l.client.Conn.IsClosed() {
		return fmt.Errorf("no rmq connection")
	}

	ch, err := l.client.Conn.Channel()
	if err != nil {
		return err
	}
	l.channel = ch

	err = ch.Qos(
		l.cfg.PrefetchCount,
		0,
		false,
	)
	if err != nil {
		return err
	}

	queueDeclareArgs := make(amqp.Table)

	queue, err := ch.QueueDeclarePassive(
		l.cfg.QueueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		queueDeclareArgs,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}
	l.msgs = msgs

	l.reconnectionError = l.client.NotifyReconnection(make(chan error, 1))
	l.closeError = l.channel.NotifyClose(make(chan *amqp.Error, 1))

	l.client.Logger.Info(fmt.Sprintf("listener %s connected", l.cfg.QueueName))

	return nil
}

func (l *Listener) reconnect() error {
	l.client.Logger.Info(fmt.Sprintf("listener %s reconnecting...", l.cfg.QueueName))

	err := <-l.reconnectionError
	if err != nil {
		return err
	}

	err = l.connect()

	return err
}

// Run it starts queue listener
func (l *Listener) Run() error {
	defer func() {
		if l.channel != nil {
			l.channel.Close()
		}
	}()
	err := l.connect()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for {
		select {
		case msg, ok := <-l.msgs:
			if ok {
				wg.Add(1)
				go func(msg amqp.Delivery) {
					var err error
					if l.hasProcessWithRoutingKeyFunc {
						if l.processWithRoutingKeyFunc != nil {
							err = l.processWithRoutingKeyFunc(msg.RoutingKey, msg.Body)
						}
					} else {
						if l.processor != nil {
							err = l.processor.Process(msg.Body)
						}
					}
					if err != nil {
						err = msg.Reject(false)
					} else {
						err = msg.Ack(false)
					}
					if err != nil {
						l.client.Logger.Error(err)
					}

					wg.Done()
				}(msg)
			} else {
				reconnErr := l.reconnect()
				if reconnErr != nil {
					return reconnErr
				}
			}
		case <-l.quit:
			wg.Wait()
			l.client.Logger.Info("all messages processed")
			return nil
		case <-l.closeError:
			reconnErr := l.reconnect()
			if reconnErr != nil {
				return reconnErr
			}
		}
	}
}

func (l *Listener) Close() {
	l.quit <- struct{}{}
}
