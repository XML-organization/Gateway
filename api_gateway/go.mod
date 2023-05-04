module api_gateway

go 1.20

replace github.com/XML-organization/common => ../common

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/XML-organization/common v1.0.1-0.20230504130318-3a270381458d // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)
