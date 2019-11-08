package main

import (
	"github.com/micro/go-config"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"shopping/notification/subscriber"
)

func main() {

	err := config.LoadFile("./config.json")
	if err != nil{
		log.Fatalf("Could not load config file: %s",err.Error())
	}
	conf := config.Map()

	b := rabbitmq.NewBroker(
		broker.Addrs(conf["rabbitmq_addr"].(string)),
	)

	b.Init()
	b.Connect()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.notification"),
		micro.Version("latest"),
		micro.Broker(b),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//notification.RegisterNotificationHandler(service.Server(), new(handler.Notification))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("notification.submit", service.Server(), new(subscriber.Notification))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.notification", service.Server(), subscriber.Handler)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
