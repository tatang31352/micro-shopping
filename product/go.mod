module shopping/product

go 1.13

require (
	github.com/go-log/log v0.1.0
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-micro v1.14.0
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6 // indirect
	google.golang.org/grpc v1.25.0 // indirect
)

replace (
	github.com/micro/go-micro => github.com/micro/go-micro v1.13.2
	google.golang.org/grpc => google.golang.org/grpc v1.19.0
)
