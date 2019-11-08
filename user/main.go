package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"shopping/user/handler"
	"shopping/user/model"
	user "shopping/user/proto/user"
	"shopping/user/repository"
)


func main() {

	err := config.LoadFile("./config.json")
	if err != nil{
		log.Fatalf("Could not load config file: %s",err.Error())
	}
	conf := config.Map()

	db,err := CreateConnection(conf["mysql"].(map[string]interface{}))
	defer db.Close()

	db.AutoMigrate(&model.User{})

	if err != nil{
		log.Fatalf("connection error : %v \n",err)
	}

	repo := &repository.User{db}
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	//repo := &repository.User{db}
	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserServiceHandler(service.Server(), &handler.User{repo})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
