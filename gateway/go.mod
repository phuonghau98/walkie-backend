module github.com/phuonghau98/walkie3/gateway

go 1.13

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	google.golang.org/grpc v1.25.1
)

require (
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/phuonghau98/walkie3/srv/user v0.0.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5
)

replace github.com/phuonghau98/walkie3/srv/user => ../srv/user

require github.com/phuonghau98/walkie3/srv/geolocation v0.0.0

replace github.com/phuonghau98/walkie3/srv/geolocation => ../srv/geolocation

require github.com/phuonghau98/walkie3/srv/message v0.0.0

replace github.com/phuonghau98/walkie3/srv/message => ../srv/message
