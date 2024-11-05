package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcServer.RegisterService(&proto.Mss, &server{})
	proto.RegisterMessageServiceServer(grpcServer, &server{})
	log.Println("Message Service is running on port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

