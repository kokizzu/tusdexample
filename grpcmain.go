package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/kokizzu/gotro/L"
	"google.golang.org/grpc"
	"log"
	"net"
	"tusdexample/tusdhooks"
)

type TusdHooksServer struct {
	tusdhooks.UnimplementedHookServiceServer
}

func (t TusdHooksServer) Send(ctx context.Context, request *tusdhooks.SendRequest) (*tusdhooks.SendResponse, error) {
	L.Describe(request.Hook)
	return &tusdhooks.SendResponse{Response: &any.Any{
		Value: []byte(`{"status":"ok"}`),
	}}, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8083")
	if err != nil {
		log.Fatalf("[ERROR] Failed to listen tcp: %v", err)
	}

	grpcServer := grpc.NewServer()
	tusdHandlers := TusdHooksServer{}
	tusdhooks.RegisterHookServiceServer(grpcServer, tusdHandlers)

	log.Println("gRPC server starting...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Println("gRPC server is running...")

}
