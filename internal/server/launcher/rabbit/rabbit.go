package rabbit

import (
	"context"
	"fmt"
	"github.com/op/go-logging"
	"golang.org/x/sync/errgroup"
	"notify-service/internal/broker/rabbit"
	"notify-service/internal/server/launcher"
	"notify-service/pkg/msgHandler"
	"sync"
)

var log = logging.MustGetLogger("rabbit")

type server struct {
	handler msgHandler.MsgHandler
	broker  rabbit.Service
}

func New(broker rabbit.Service, handler msgHandler.MsgHandler) launcher.Server {
	return &server{
		handler: handler,
		broker:  broker,
	}
}

func (s server) Serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(1)

	gr, grCtx := errgroup.WithContext(ctx)

	gr.Go(func() error {
		defer wg.Done()
		return s.serve(grCtx)
	})

	wg.Wait()

	return nil
}

func (s server) serve(ctx context.Context) error {
	c := s.broker.Consumer()

	for {
		if err := ctx.Err(); err != nil {
			log.Criticalf("Rabbit listener stopped error: %v", err)
			return fmt.Errorf("rabbit listener stopped error: %v", err)
		}

		m, err := c.Consume(ctx)
		if err != nil {
			log.Infof("failed to consume message error: %v", err)
			continue
		}

		go func() {
			err := s.handler.ServeMSG(ctx, m)
			if err != nil {
				log.Criticalf("failed to handle message: %v", err)
				return
			}
		}()
	}
}
