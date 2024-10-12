package rabbit

import amqp "github.com/rabbitmq/amqp091-go"

func Configure(ch *amqp.Channel) error {

	err := ch.ExchangeDeclare(
		"test",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"yellow", // name
		false,    // durable
		false,    // delete when unused
		true,     // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,   // name
		"yellow", // key
		"test",   // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}
