module shopping/notification

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-config v1.1.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.1.0
	product v0.0.0
	user v0.0.0
)

replace (
	product => ../product
	user => ../user
)

replace (
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422
	github.com/hashicorp/consul v1.4.3 => github.com/hashicorp/consul v1.4.3
	github.com/micro/go-micro => github.com/micro/go-micro v1.1.0
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.3
)
