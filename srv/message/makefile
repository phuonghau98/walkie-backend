build:
	protoc -I. --go_out=plugins=grpc:. proto/message/message.proto
	protoc -I/usr/local/include -I. \
		--grpc-gateway_out=logtostderr=true,grpc_api_configuration=proto/message/message_service.yaml:. \
		proto/message/message.proto