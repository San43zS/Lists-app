package consumer

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	Consumer(ctx context.Context) ([]byte, error)
	Close() error
}

type consumer struct {
	dial *amqp.Channel
}

func New(dial *amqp.Channel) Consumer {
	return &consumer{
		dial: dial,
	}
}

func (c consumer) Consumer(ctx context.Context) ([]byte, error) {

	return nil, nil
}

func (c consumer) Close() error {
	return nil
}
