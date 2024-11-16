package event

import (
	msg2 "Lists-app/internal/handler/model/msg"
	"context"
	"fmt"
	"log"
)

func (h handler) AddNotify(ctx context.Context, msg msg2.MSG) error {
	err := h.respondConsumer.p.Produce(ctx, msg.Content.Data)
	if err != nil {
		log.Println("Failed to add notification: ", err)
		return fmt.Errorf("failed to add notification: %w", err)
	}

	return nil
}

func (h handler) GetCurrentNotify(ctx context.Context) ([]byte, error) {
	consume, err := h.respondConsumer.c.Consume(ctx)
	if err != nil {
		log.Println("Failed to get notifications with ttl: ", err)
		return nil, fmt.Errorf("failed to get notifications with ttl: %w", err)
	}

	return consume, nil
}

func (h handler) GetOldNotify(ctx context.Context) ([]byte, error) {
	consume, err := h.respondConsumer.c.Consume(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to get notifications without ttl:  %w", err)
		log.Println(err.Error())
		
		return nil, err
	}

	return consume, nil
}
