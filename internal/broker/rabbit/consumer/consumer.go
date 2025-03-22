package consumer

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"notify-service/internal/broker/rabbit/config"
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

	msgs, err := c.dial.Consume(
		config.ConsumerQueueName, // queue
		"",                       // consumer
		false,                    // auto-ack
		false,                    // exclusive
		false,                    // no-local
		true,                     // no-wait
		nil,                      // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume message: %w", err)
	}

	for msg := range msgs {
		return msg.Body, nil
	}

	return nil, nil
}
