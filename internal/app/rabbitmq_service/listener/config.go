package listener

import "time"

type AMQPConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Hostname string `yaml:"hostname"`
	Protocol string `yaml:"protocol"`
	VHost    string `yaml:"vhost"`
	Port     int32  `yaml:"port"`

	ReconnectTimeout time.Duration `yaml:"reconnect_timeout_ms"`
}

type ListenerConfig struct {
	QueueName     string `yaml:"queue_name"`
	PrefetchCount int    `yaml:"prefetch_count"`
}
