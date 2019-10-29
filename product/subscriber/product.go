package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	product "shopping/product/proto/product"
)

type Product struct{}

func (e *Product) Handle(ctx context.Context, msg *product.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *product.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
