package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	healer "healer/proto/healer"
)

type Healer struct{}

func (e *Healer) Handle(ctx context.Context, msg *healer.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *healer.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
