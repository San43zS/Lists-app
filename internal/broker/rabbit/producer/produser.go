package producer

import (
	"Lists-app/internal/broker/rabbit/config"
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer interface {
	Produce(ctx context.Context, arr []byte) error
}

type producer struct {
	dial *amqp.Channel
}

func New(dial *amqp.Channel) Producer {
	return &producer{
		dial: dial,
	}
}

func (p producer) Produce(ctx context.Context, msg []byte) error {
	f := string(msg)

	err := p.dial.PublishWithContext(
		ctx,
		config.ProducerExchangeName, // exchange
		config.ProducerRoutingKey,   // routing key
		false,                       // mandatory
		false,                       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(f),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
