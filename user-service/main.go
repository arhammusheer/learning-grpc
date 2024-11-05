package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    "github.com/arhammusheer/learning-grpc/proto/user"
)

type userServiceServer struct {
    user.UnimplementedUserServiceServer
}

func NewUserServiceServer() *userServiceServer {
    return &userServiceServer{}
}

func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    user.RegisterUserServiceServer(grpcServer, NewUserServiceServer())

    log.Println("User Service is running on port 50051...")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
