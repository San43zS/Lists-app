package rabbitMQ

import (
	"context"
	"log"
	"time"

	"Lists-app/internal/service/rabbitMQ/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RbbMQ interface {
	Connect(config config.Config) error
	Send() error
}

type RabbitMQ struct {
	rbbMQ RbbMQ
}

func (r RabbitMQ) Connect(config config.Config) error {
	conn, err := amqp.Dial(config.Driver + config.URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.rbbMQ.Connect()
}

func (r RabbitMQ) Send() error {
	return r.rbbMQ.Send()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
