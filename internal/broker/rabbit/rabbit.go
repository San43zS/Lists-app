package rabbit

import (
	"Lists-app/internal/broker/rabbit/config"
	"Lists-app/internal/broker/rabbit/consumer"
	"Lists-app/internal/broker/rabbit/producer"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Service interface {
	Producer() producer.Producer
	Consumer() consumer.Consumer
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

	srv := &service{dial: ch}

	srv.Producer().Produce()

	return srv, nil
}

func (s service) Producer() producer.Producer {
	return producer.New(s.dial)
}

func (s service) Consumer() consumer.Consumer {
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
