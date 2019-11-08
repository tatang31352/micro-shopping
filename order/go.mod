module shopping/order

go 1.13

require (
	github.com/bwmarrin/snowflake v0.0.0-20180412010544-68117e6bbede
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/go-log/log v0.1.0
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/micro/go-config v1.1.0
	github.com/micro/go-grpc v1.0.1 // indirect
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.1.0
	github.com/prometheus/client_golang v1.1.0
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	shopping/product v0.0.0
)

replace (
	github.com/golang/protobuf v1.3.2 => github.com/golang/protobuf v1.3.1
	github.com/jinzhu/gorm v1.9.11 => github.com/jinzhu/gorm v1.9.2
	github.com/micro/go-micro => github.com/micro/go-micro v1.1.0
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
)

replace (
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.3
	shopping/product => ../product
)
