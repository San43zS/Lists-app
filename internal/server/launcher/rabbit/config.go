package rabbit

import "notify-service/internal/broker/rabbit/config"

type consumer struct {
	topic string
	ID    string
}

type Model struct {
	consumers []consumer
}

func NewModel() Model {
	model := Model{}

	model.consumers = append(model.consumers, consumer{
		topic: config.ProducerQueueName,
		ID:    config.ProducerQueueName,
	})

	return model
}
