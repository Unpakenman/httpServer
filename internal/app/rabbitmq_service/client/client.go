package client

import (
	"context"
	"fmt"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Logger  LogClient
	connURL string
	cfg     AMQPConfig
	quit    chan bool

	Conn       *amqp.Connection
	closeError chan *amqp.Error

	reconnectionListenersMu sync.Mutex
	reconnectionListeners   []chan error
}

//go:generate mockery --with-expecter --case=underscore --name=LogClient

type LogClient interface {
	Info(msg string, fields ...interface{})
	InfoCtx(ctx context.Context, msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(err error, fields ...interface{})
}

// NewClient creates a Client instance
func NewClient(cfg *AMQPConfig, logger LogClient, quit chan bool) (*Client, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid amqp config")
	}

	connectionURL := fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s",
		cfg.Protocol,
		cfg.User,
		cfg.Password,
		cfg.Hostname,
		cfg.Port,
		cfg.VHost,
	)

	c := Client{
		Logger:  logger,
		connURL: connectionURL,
		cfg:     *cfg,
		quit:    quit,
	}

	err := c.connect()
	if err != nil {
		return nil, err
	}

	go c.watchConnection()

	return &c, nil
}

func (c *Client) connect() error {
	if c.Conn != nil && !c.Conn.IsClosed() {
		return nil
	}

	conn, err := amqp.Dial(c.connURL)
	if err != nil {
		return err
	}
	c.Conn = conn

	c.closeError = conn.NotifyClose(make(chan *amqp.Error))

	return nil
}

func (c *Client) reconnect() {
	for {
		err := c.connect()
		if err != nil {
			c.Logger.Warn("trying reconnect to RabbitMQ error", err.Error())

			time.Sleep(c.cfg.ReconnectTimeout)
			continue
		}

		c.reconnectionListenersMu.Lock()
		for _, l := range c.reconnectionListeners {
			close(l)
		}

		c.reconnectionListeners = nil
		c.reconnectionListenersMu.Unlock()

		return
	}
}

func (c *Client) watchConnection() {
	defer func() {
		for _, l := range c.reconnectionListeners {
			close(l)
		}
		if c.Conn != nil {
			c.Conn.Close()
		}
	}()

	for {
		select {
		case <-c.closeError:
			c.reconnect()
		case <-c.quit:
			return
		}
	}
}

func (c *Client) NotifyReconnection(listener chan error) chan error {
	c.reconnectionListenersMu.Lock()
	defer c.reconnectionListenersMu.Unlock()

	c.reconnectionListeners = append(c.reconnectionListeners, listener)

	return listener
}

func (c *Client) GetConnection() *amqp.Connection {
	return c.Conn
}
