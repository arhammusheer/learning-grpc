package main

import (
    "context"
    "sync"
    "log"

    "github.com/arhammusheer/learning-grpc/proto/user"
)

type userServiceServer struct {
    user.UnimplementedUserServiceServer
    users map[string]*user.UserResponse
    mu    sync.Mutex // For safe concurrent access
}

// NewUserServiceServer creates a new UserServiceServer instance
func NewUserServiceServer() *userServiceServer {
    return &userServiceServer{
        users: make(map[string]*user.UserResponse),
    }
}

// CreateUser adds a new user to the in-memory store
func (s *userServiceServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Check if user already exists
    if _, exists := s.users[req.UserId]; exists {
        log.Printf("User %s already exists", req.UserId)
        return nil, status.Errorf(codes.AlreadyExists, "User %s already exists", req.UserId)
    }

    // Create and store the user
    newUser := &user.UserResponse{
        UserId:   req.UserId,
        Name:     req.Name,
        Email:    req.Email,
    }
    s.users[req.UserId] = newUser

    log.Printf("User %s created", req.UserId)
    return newUser, nil
}

// GetUser retrieves a user by user ID
func (s *userServiceServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    user, exists := s.users[req.UserId]
    if !exists {
        log.Printf("User %s not found", req.UserId)
        return nil, status.Errorf(codes.NotFound, "User %s not found", req.UserId)
    }

    return user, nil
}
