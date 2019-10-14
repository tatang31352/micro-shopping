package main

import (
	"demo/micro/shopping/user/handler"
	"demo/micro/shopping/user/model"
	user "demo/micro/shopping/user/proto/user"
	"demo/micro/shopping/user/repository"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/service/grpc"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	//加载配置项
	err := config.LoadFile("./config.json")
	if err != nil{
		log.Fatal("Could not load config file: %s",err.Error())
		return
	}
	conf := config.Map()

	//db
	db,err := CreateConnection(conf["mysql"].(map[string]interface{}))
	defer db.Close()

	db.AutoMigrate(&model.User{})

	if err != nil{
		log.Fatalf("connection error : %v \n",err)
	}

	repo := &repository.User{db}

	//New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	//Register Handler
	user.RegisterUserServiceHandler(service.Server(), &handler.User{repo})

	//Register Struct ad Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user",service.Server(),new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	//Run service
	if err := service.Run();err != nil{
		log.Fatal(err)
	}

}


func CreateConnection(conf map[string]interface{})(*gorm.DB,error){
	host := conf["host"]
	port := conf["port"]
	user := conf["user"]
	dbName := conf["database"]
	password := conf["password"]
	return gorm.Open("mysql",fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",user,password,host,port,dbName,
	),
	)
}
