build:
	protoc -I. --go_out=plugins=grpc:. proto/user/user.proto
	protoc -I/usr/local/include -I. \
    --grpc-gateway_out=logtostderr=true,request_context=true,grpc_api_configuration=proto/user/user_service.yaml:. \
    proto/user/user.proto
#   protoc -I/usr/local/include -I. \
#        --swagger_out=logtostderr=true,grpc_api_configuration=proto/user/user_service.yaml:. \
#        proto/user/user.proto