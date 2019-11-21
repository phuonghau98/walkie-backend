package main

import (
	"github.com/phuonghau98/walkie3/srv/user/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	//"golang.org/x/crypto/bcrypt"
)

var (
	key = []byte("mySuperSecretKeyLol")
)

type Service struct {
	repo *UserRepository
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	dbConnection, err := CreateConnection()
	dbConnection.AutoMigrate(&User{})
	defer dbConnection.Close()
	if err != nil {
		log.Fatalf("cannot connect to database")
	}

	repo := &UserRepository{db: dbConnection}

	user.RegisterUserServiceServer(s, &Service{repo})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
