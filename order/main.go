package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"shopping/order/handler"
	"shopping/order/model"
	order "shopping/order/proto/order"
	"shopping/order/repository"
	product "shopping/product/proto/product"
)

func main() {

	db,err := CreateConnection()
	defer db.Close()

	db.AutoMigrate(&model.Order{})

	if err != nil{
		log.Fatal("connection error : %v \n",err)
	}

	repo := &repository.Order{db}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// 创建消息发布者
	publisher := micro.NewPublisher("notification.submit",service.Client())

	//product-srv client
	productCli := product.NewProductService("go.micro.srv.product",service.Client())

	//Registart Handler
	order.RegisterOrderServiceHandler(service.Server(),&handler.Order{repo,productCli,publisher})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.order", service.Server(), new(subscriber.Order))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.order", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
