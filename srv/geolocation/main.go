package main

import (
	"context"
	"github.com/phuonghau98/walkie3/srv/geolocation/proto/geolocation"
	"github.com/phuonghau98/walkie3/srv/message/proto/message"
	"github.com/phuonghau98/walkie3/srv/user/proto/user"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"log"
	"net"
)
var (
	userServiceAddress = "localhost:8081"
	messageServiceAddress = "localhost:8083"
)

type Service struct {
	repo *GeolocationRepository
	userClient user.UserServiceClient
	messageClient message.MessageServiceClient
}

func main() {
	listener, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userServiceConnection, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	messageServiceConnection, err := grpc.Dial(messageServiceAddress, grpc.WithInsecure())
	defer userServiceConnection.Close()

	s := grpc.NewServer()
	dbConnection := createConnection()
	defer dbConnection.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("cannot connect to database")
	}

	repo := &GeolocationRepository{dbConnection.Database("walkie3")}

	geolocation.RegisterGeolocationServiceServer(s, &Service{
		repo,
		user.NewUserServiceClient(userServiceConnection),
		message.NewMessageServiceClient(messageServiceConnection),
	})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
