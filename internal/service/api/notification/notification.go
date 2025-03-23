package notification

import (
	"context"
)

type Notification interface {
	Add(ctx context.Context, msg msg2.MSG) error

	GetOld(ctx context.Context) ([]byte, error)

	GetCurrent(ctx context.Context) ([]byte, error)
}
