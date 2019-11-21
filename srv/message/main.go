package main

import (
	"context"
	"github.com/phuonghau98/walkie3/srv/message/proto/message"
	"github.com/phuonghau98/walkie3/srv/user/proto/user"

	//"github.com/phuonghau98/walkie3/srv/user/proto/user"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"log"
	"net"
)
var (
	userServiceAddress = "localhost:8081"
)

type Service struct {
	repo *MessageRepository
	userClient user.UserServiceClient
}

func main() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userServiceConnection, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	defer userServiceConnection.Close()

	s := grpc.NewServer()
	dbConnection := createConnection()
	defer dbConnection.Disconnect(context.Background())
	//dbConnection.AutoMigrate(&User{})
	//defer dbConnection.Close()
	if err != nil {
		log.Fatalf("cannot connect to database")
	}

	repo := &MessageRepository{ dbConnection.Database("walkie3")}

	message.RegisterMessageServiceServer(s, &Service{repo, user.NewUserServiceClient(userServiceConnection)})

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
