package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	order "shopping/order/proto/order"
)

type Order struct{}

func (e *Order) Handle(ctx context.Context, msg *order.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *order.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
