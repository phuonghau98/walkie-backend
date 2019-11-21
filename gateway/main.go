package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	messagePb "github.com/phuonghau98/walkie3/srv/message/proto/message"
	geoPb "github.com/phuonghau98/walkie3/srv/geolocation/proto/geolocation"
	"net/http"

	userPb "github.com/phuonghau98/walkie3/srv/user/proto/user"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	userGRPCServer = flag.String("user-grpc-server", "localhost:8081", "gRPC server endpoint")
	geolocationGRPCServer = flag.String("geolocation-grpc-server", "localhost:8084", "gRPC server endpoint")
	messageGRPCServer = flag.String("message-grpc-server", "localhost:8083", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := userPb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *userGRPCServer, opts)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	err = geoPb.RegisterGeolocationServiceHandlerFromEndpoint(ctx, mux, *geolocationGRPCServer, opts)
	if err != nil {
		return err
	}
	err = messagePb.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, *messageGRPCServer, opts)
	//Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8080", AllowCors(wsproxy.WebsocketProxy(mux)))
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
