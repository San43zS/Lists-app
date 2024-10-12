package rabbitMQ

import (
	"Lists-app/internal/broker/rabbitMQ/config"
	"Lists-app/internal/broker/rabbitMQ/consumer"
	"Lists-app/internal/broker/rabbitMQ/producer"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Service interface {
	Produce() producer.Producer
	Consume() consumer.Consumer
}

type service struct {
	dial *amqp.Channel
}

func New() (Service, error) {
	cfg := config.NewConfig()

	conn, err := amqp.Dial(cfg.Driver + cfg.URL)
	if err != nil {
		conn.Close()
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {

		ch.Close()
		conn.Close()
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"test",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &service{dial: ch}, nil
}

func (s service) Produce() producer.Producer {
	return producer.New(s.dial)
}

func (s service) Consume() consumer.Consumer {
	return consumer.New(s.dial)
}

func Send(config config.Config, msg string) {
	conn, err := amqp.Dial(config.Driver + config.URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"Rabbit", // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := msg
	err = ch.PublishWithContext(ctx,
		"Rabbit", // exchange
		"yellow", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
