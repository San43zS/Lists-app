package producer

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Producer interface {
	Produce(ctx context.Context, arr []byte) error
	Close() error
}

type producer struct {
	dial *amqp.Channel
}

func New(dial *amqp.Channel) Producer {
	return &producer{
		dial: dial,
	}
}

func (p producer) Produce(ctx context.Context, arr []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := arr
	err := p.dial.PublishWithContext(ctx,
		"test",   // exchange
		"yellow", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		return err
	}

	return nil
}

func (p producer) Close() error {
	p.dial.Close()
	return nil
}
