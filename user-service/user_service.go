package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"user/proto"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	users map[string]*proto.UserResponse
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		users: make(map[string]*proto.UserResponse),
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	userID := uuid.New().String()
	user := &proto.UserResponse{
		UserId: userID,
		Name:   req.Name,
		Email:  req.Email,
	}
	s.users[userID] = user
	return user, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	user, exists := s.users[req.UserId]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
