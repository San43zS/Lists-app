package broker

import (
	"Lists-app/internal/broker/rabbitMQ"
	"fmt"
)

type Broker struct {
	RabbitMQ rabbitMQ.Service
}

func New() (Broker, error) {
	rabbitMQ, err := rabbitMQ.New()
	if err != nil {
		return Broker{}, fmt.Errorf("failed to create RabbitMQ broker: %w", err)
	}

	broker := Broker{
		RabbitMQ: rabbitMQ,
	}

	return broker, nil
}
