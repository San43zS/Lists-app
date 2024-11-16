package event

import (
	"Lists-app/internal/model/notification"
	"context"
)

func (h handler) AddNotify(ctx context.Context, msg []byte) error {
	n := notification.Notification{
		Id:        1,
		Info:      "string(msg)",
		TTL:       10,
		CreatedAt: 1,
	}
	h.srv.Notification().Add(ctx, n)
	return nil
}

func (h handler) GetNotify(ctx context.Context, msg []byte) error {
	return nil
}
