package notification

import (
	"context"
	"fmt"
	"github.com/op/go-logging"
	"notify-service/internal/broker"
	"notify-service/internal/broker/rabbit/consumer"
	"notify-service/internal/broker/rabbit/producer"
)

var log = logging.MustGetLogger("service")

type RespCons struct {
	p producer.Producer
	c consumer.Consumer
}

func New(broker broker.Broker) RespCons {
	return RespCons{
		p: broker.RabbitMQ.Producer(),
		c: broker.RabbitMQ.Consumer(),
	}
}

func (s RespCons) Add(ctx context.Context, msg msg2.MSG) error {
	newMsg, err := msg2.New().Unparse(msg)
	if err != nil {
		log.Criticalf("Failed to add notification: %w", err)
		return fmt.Errorf("failed to add notification: %w", err)
	}

	err = s.p.Produce(ctx, newMsg)
	if err != nil {
		log.Criticalf("Failed to add notification: ", err)
		return fmt.Errorf("failed to add notification: %w", err)
	}

	return nil
}

func (s RespCons) GetOld(ctx context.Context) ([]byte, error) {
	consume, err := s.c.Consume(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to get notifications without ttl:  %w", err)
		log.Criticalf(err.Error())

		return nil, err
	}

	return consume, nil
}

func (s RespCons) GetCurrent(ctx context.Context) ([]byte, error) {
	consume, err := s.c.Consume(ctx)
	if err != nil {
		log.Criticalf("Failed to get notifications with ttl: ", err)
		return nil, fmt.Errorf("failed to get notifications with ttl: %w", err)
	}

	return consume, nil
}
