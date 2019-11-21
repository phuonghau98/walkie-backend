build:
	protoc -I. --go_out=plugins=grpc:. proto/geolocation/geolocation.proto
	protoc -I/usr/local/include -I. \
    --grpc-gateway_out=logtostderr=true,grpc_api_configuration=proto/geolocation/geolocation_service.yaml:. \
    proto/geolocation/geolocation.proto
  protoc -I/usr/local/include -I. \
       --swagger_out=logtostderr=true,grpc_api_configuration=proto/geolocation/geolocation_service.yaml:. \
       proto/geolocation/geolocation.proto