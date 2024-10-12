package consumer

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Consumer interface {
	Consume(ctx context.Context) ([]byte, error)
}

type consumer struct {
	dial *amqp.Channel
}

func New(dial *amqp.Channel) Consumer {
	return &consumer{
		dial: dial,
	}
}

func (c consumer) Consume(ctx context.Context) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgs, err := c.dial.Consume(
		"yellow", // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return nil, err
	}

	for msg := range msgs {
		return msg.Body, nil
	}

	return nil, nil
}
