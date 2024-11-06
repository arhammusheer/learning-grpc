package main

import (
	"context"

	pb "arhammusheer/learning-grpc/proto/user"
)

// NewUserServiceServer creates a new UserServiceServer instance
type server struct {
	pb.UnimplementedUserServiceServer
}

// CreateUser creates a new user
func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		UserId: "1",
		Name:   req.GetName(),
		Email:  req.GetEmail(),
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		UserId: req.GetUserId(),
		Name:   "Arham Musheer",
		Email:  "arham@croissant.one",
	}, nil
}
