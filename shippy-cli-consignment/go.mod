module github.com/pallymore/shippy/shippy-cli-consignment

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pallymore/shippy/shippy-service-consignment v0.0.0-20200712014547-abdb0c48f959
)
