module shopping/order

go 1.13

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/go-log/log v0.1.0
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-micro v1.14.0
	shopping/product v0.0.0
)

replace (
	github.com/micro/go-micro => github.com/micro/go-micro v1.13.2
	google.golang.org/grpc => google.golang.org/grpc v1.19.0
	shopping/product => ../product
)
