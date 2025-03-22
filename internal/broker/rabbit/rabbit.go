package rabbit

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"notify-service/internal/broker/rabbit/config"
	"notify-service/internal/broker/rabbit/consumer"
	"notify-service/internal/broker/rabbit/producer"
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
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	if err := ConfigureConsumer(ch); err != nil {
		ch.Close()
		conn.Close()

		return nil, err
	}

	if err := ConfigureProducer(ch); err != nil {
		ch.Close()
		conn.Close()

		return nil, err
	}

	srv := &service{dial: ch}

	return srv, nil
}

func (s service) Producer() producer.Producer {
	return producer.New(s.dial)
}

func (s service) Consumer() consumer.Consumer {
	return consumer.New(s.dial)
}
