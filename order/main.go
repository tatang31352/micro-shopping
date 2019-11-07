package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-log"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"shopping/order/handler"
	"shopping/order/model"
	order "shopping/order/proto/order"
	"shopping/order/repository"
	product "shopping/product/proto/product"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	wrapperPrometheus "github.com/micro/go-plugins/wrapper/monitoring/prometheus"
)

func main() {

	db,err := CreateConnection()
	defer db.Close()

	db.AutoMigrate(&model.Order{})

	if err != nil{
		log.Fatal("connection error : %v \n",err)
	}

	repo := &repository.Order{db}

	//broker
	b := rabbitmq.NewBroker(
		broker.Addrs("amqp://guest:guest@127.0.0.1:5672"),
	)

	//boot trace
	TraceBoot()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.Version("latest"),
		micro.Broker(b),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper()),
		micro.WrapClient(wrapperTrace.NewClientWrapper()),
		micro.WrapHandler(wrapperPrometheus.NewHandlerWrapper()),
	)

	//boot prometheus
	PrometheusBoot()

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


func TraceBoot(){

	serviceName := "go.micro.srv.order"

	cfg := jaegercfg.Configuration{
		Sampler:&jaegercfg.SamplerConfig{
			Type:"const",
			Param:1,
		},
		Reporter:&jaegercfg.ReporterConfig{
			LogSpans:true,
			LocalAgentHostPort:"127.0.0.1:9412",
		},
	}

	closer,err := cfg.InitGlobalTracer(
		serviceName,
	)
	if err != nil{
		log.Fatalf("Could not initialize jaeger tracer: %s",err.Error())
		return
	}
	defer closer.Close()

	//opentracing.InitGlobalTracer(tracer)
	return
}

func PrometheusBoot(){
	http.Handle("/metrics",promhttp.Handler())
	//启动web服务,监听8085端口
	go func() {
		err := http.ListenAndServe("127.0.0.1:8085",nil)
		if err != nil{
			log.Fatal("ListenAndServe: ",err)
		}
	}()
}


