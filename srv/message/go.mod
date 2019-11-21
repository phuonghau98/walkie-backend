module github.com/phuonghau98/walkie3/srv/message

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/tidwall/pretty v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.3
	google.golang.org/grpc v1.25.1
)

require (
	github.com/phuonghau98/walkie3/srv/user v0.0.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
)

replace github.com/phuonghau98/walkie3/srv/user => ../user

replace github.com/phuonghau98/walkie3/srv/message => ../message
