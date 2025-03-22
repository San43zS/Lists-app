package event

import (
	"context"
	"encoding/json"
	"fmt"
	message "notify-service/internal/handler/model/msg"
)

var MasCurrent message.STRUCT
var MasOld message.STRUCT

func (h handler) ShowCurrent(ctx context.Context, msg []byte) error {
	var temp message.STRUCT
	err := json.Unmarshal(msg, &temp)
	if err != nil {
		return fmt.Errorf("error while parsing(unmarshal) msg: %w", err)
	}

	MasCurrent = temp
	return nil
}

func (h handler) ShowOld(ctx context.Context, msg []byte) error {
	var temp message.STRUCT
	err := json.Unmarshal(msg, &temp)
	if err != nil {
		return fmt.Errorf("error while parsing(unmarshal) msg: %w", err)
	}

	MasOld = temp
	return nil
}
