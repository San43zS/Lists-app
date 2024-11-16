package notification

import (
	msg2 "Lists-app/internal/handler/model/msg"
	"context"
)

type Notification interface {
	Add(ctx context.Context, msg msg2.MSG) error

	GetOld(ctx context.Context) ([]byte, error)

	GetCurrent(ctx context.Context) ([]byte, error)
}
