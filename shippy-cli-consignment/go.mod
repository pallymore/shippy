module github.com/pallymore/shippy/shippy-cli-consignment

go 1.14

require (
	github.com/pallymore/shippy/shippy-service-consignment v0.0.0-20200712004859-b595e4dd0287
	google.golang.org/grpc v1.30.0
)

replace github.com/pallymore/shippy/shippy-service-consignment => ../shippy-service-consignment
