package main

import (
	pb "arhammusheer/learning-grpc/proto/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the user service with the server
	pb.RegisterUserServiceServer(grpcServer, &server{})
	log.Println("User Service is running on port 50051...")

	// Enable reflection for gRPC server
	reflection.Register(grpcServer)

	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
