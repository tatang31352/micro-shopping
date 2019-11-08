package main

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"shopping/product/handler"
	"shopping/product/model"
	product "shopping/product/proto/product"
	"shopping/product/repository"
)

func main() {

	err := config.LoadFile("./config.json")
	if err != nil{
		log.Fatalf("Could not load config file: %s",err.Error())
	}
	conf := config.Map()

	db,err := CreateConnection(conf["mysql"].(map[string]interface{}))
	defer db.Close()

	db.AutoMigrate(&model.Product{})

	if err != nil{
		log.Fatalf("connect error : %v \n",err)
	}

	repo := &repository.Product{db}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.product"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductServiceHandler(service.Server(), &handler.Product{repo})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.product", service.Server(), new(subscriber.Product))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.product", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
