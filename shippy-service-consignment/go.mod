module github.com/pallymore/shippy/shippy-service-consignment

go 1.14

// TODO: remove this in prod
replace github.com/pallymore/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1 // indirect
	github.com/micro/micro/v2 v2.9.2-0.20200711221010-7f17a2208222 // indirect
	github.com/spf13/viper v1.6.3 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200711021454-869866162049 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)
