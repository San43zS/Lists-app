package rabbit

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"notify-service/internal/broker/rabbit/config"
)

func ConfigureConsumer(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		config.ConsumerExchangeName, // name
		amqp.ExchangeDirect,         // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	q, err := ch.QueueDeclare(
		config.ConsumerQueueName, // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.QueueBind(
		q.Name,                      // name
		config.ConsumerRoutingKey,   // key
		config.ConsumerExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}
	return nil
}

func ConfigureProducer(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		config.ProducerExchangeName, // name
		amqp.ExchangeDirect,         // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	q, err := ch.QueueDeclare(
		config.ProducerQueueName, // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.QueueBind(
		q.Name,                      // name
		config.ProducerRoutingKey,   // key
		config.ProducerExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	return nil
}
