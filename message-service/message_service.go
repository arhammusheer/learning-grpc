package main

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
	"message/proto"
)

type MessageServiceServer struct {
	proto.UnimplementedMessageServiceServer
	messages map[string][]*proto.MessageResponse
}

func NewMessageServiceServer() *MessageServiceServer {
	return &MessageServiceServer{
		messages: make(map[string][]*proto.MessageResponse),
	}
}

func (s *MessageServiceServer) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.MessageResponse, error) {
	message := &proto.MessageResponse{
		MessageId: uuid.New().String(),
		UserId:    req.UserId,
		Content:   req.Content,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	s.messages[req.UserId] = append(s.messages[req.UserId], message)
	return message, nil
}

func (s *MessageServiceServer) GetMessages(ctx context.Context, req *proto.GetMessagesRequest) (*proto.MessagesResponse, error) {
	userMessages, exists := s.messages[req.UserId]
	if !exists {
		return nil, fmt.Errorf("no messages found for user")
	}
	return &proto.MessagesResponse{Messages: userMessages}, nil
}
