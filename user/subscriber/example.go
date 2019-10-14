package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"
	example "demo/micro/shopping/user/proto/order"
)

type Example struct {}

func (e *Example) Handle(ctx context.Context,msg *example.Message)error{
	log.Log("handler Received massage: ",msg.Say)
	return nil
}

func handler(ctx context.Context,msg *example.Message) error{
	log.Log("Function Received message: ",msg.Say)
	return nil
}
