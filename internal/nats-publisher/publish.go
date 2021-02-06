package natsbot

import (
	"time"

	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	Connection *nats.Conn
}

func Setup() (*NatsClient, error) {
	nc, err := nats.Connect("0.0.0.0:4222", nats.FlusherTimeout(time.Minute))
	if err != nil {
		return nil, err
	}
	return &NatsClient{connection: nc}, nil
}

// Publish write message to nats server
func (nc *NatsClient) Publish(message string) {
	nc.Connection.Publish("Test-message", []byte(message))
}
