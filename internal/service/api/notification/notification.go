package notification

import (
	"context"
	message "notify-service/internal/handler/model/msg"
)

type Notification interface {
	Add(ctx context.Context, msg message.MSG) error

	GetOld(ctx context.Context) ([]byte, error)

	GetCurrent(ctx context.Context) ([]byte, error)
}
