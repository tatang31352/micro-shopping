module shopping/user

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-micro v1.13.2
	github.com/micro/go-plugins v1.1.1
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0 // indirect
	golang.org/x/sys v0.0.0-20191105231009-c1f44814a5cd // indirect
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6 // indirect
	google.golang.org/grpc v1.25.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.19.0
