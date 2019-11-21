module github.com/phuonghau98/walkie3/srv/geolocation

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	golang.org/x/crypto v0.0.0-20191119213627-4f8c1d86b1ba // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	google.golang.org/grpc v1.25.1
)

require github.com/phuonghau98/walkie3/srv/user v0.0.0

replace github.com/phuonghau98/walkie3/srv/user => ../user

require github.com/phuonghau98/walkie3/srv/message v0.0.0

replace github.com/phuonghau98/walkie3/srv/message => ../message
