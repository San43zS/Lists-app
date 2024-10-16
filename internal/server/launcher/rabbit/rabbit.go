package rabbit

import (
	"Lists-app/internal/broker/rabbit"
	"Lists-app/internal/server/launcher"
	"Lists-app/pkg/msgHandler"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
)

type server struct {
	handler msgHandler.MsgHandler
	broker  rabbit.Service
	model   Model
}

func New(broker rabbit.Service, handler msgHandler.MsgHandler) launcher.Server {

	return &server{
		handler: handler,
		broker:  broker,
		model:   NewModel(),
	}
}

func (s server) Serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(s.model.consumers))

	gr, grCtx := errgroup.WithContext(ctx)

	for _, consumer := range s.model.consumers {
		consumer := consumer

		gr.Go(func() error {
			defer wg.Done()
			return s.serve(grCtx, consumer)
		})
	}

	wg.Wait()

	return nil
}

func (s server) serve(ctx context.Context, consumer consumer) error {
	c := s.broker.Consumer()

	for {
		if err := ctx.Err(); err != nil {
			fmt.Printf("failed to consume message [%s]: %v\n", consumer.topic, err)
			return nil
		}

		m, err := c.Consume(ctx)

		if err != nil {
			fmt.Printf("failed to consume message [%s]: %v\n", consumer.topic, err)
			continue
		}

		go func() {
			err := s.handler.ServeMSG(ctx, m)
			if err != nil {
				fmt.Printf("failed to process message [%s]: %v\n", consumer.topic, err)

				return
			}
		}()

	}
}
