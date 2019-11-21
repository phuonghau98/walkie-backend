package main

import (
	"context"
	"github.com/phuonghau98/walkie3/srv/message/proto/message"
	"github.com/phuonghau98/walkie3/srv/user/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Service) Send (ctx context.Context,  incomingMsg *message.Message) (*message.Empty, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	foundUser, err := s.userClient.GetUserFromToken(ctx, &user.TokenInfo{Token:headers["authorization"][0]})
	if err != nil {
		return nil, err
	}
	incomingMsg.Sender = foundUser.Id
	s.repo.InsertMessage(incomingMsg)
	return &message.Empty{}, nil
}

func (s *Service) Receive (ctx context.Context, req *message.MessagesRequest) (*message.Messages, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	foundUser, err := s.userClient.GetUserFromToken(ctx, &user.TokenInfo{Token:headers["authorization"][0]})
	if err != nil {
		return nil, err
	}
	var messages []*message.Message
	for _, v := range s.repo.GetMessages(foundUser.Id, req.Milestone) {
		messages = append(messages, &message.Message{
			Id: 				  v.ID.Hex(),
			Sender:               v.Sender,
			Receiver:             v.Receiver,
			Content:              v.Content,
			IsRead:               v.IsRead,
			CreatedAt:            v.CreatedAt.Unix(),
		})
	}
	return &message.Messages{Messages:messages}, nil
}

func (s *Service) ReceiveLastMessages (ctx context.Context, req *message.LastMessagesRequest) (*message.LastMessagesResponse, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	foundUser, err := s.userClient.GetUserFromToken(ctx, &user.TokenInfo{Token:headers["authorization"][0]})
	if err != nil {
		return nil, err
	}
	var messages []*message.Message
	if req.UserIDs == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Users listed is not provided")
	}
	for _, v := range req.UserIDs {
		msg, err := s.repo.GetSentMessages(foundUser.Id, v, 10)
		if err != nil {
			return nil, err
		}
		for _, v := range msg {
			messages = append(messages, &message.Message{
				Id:                   v.ID.Hex(),
				Sender:               v.Sender,
				Receiver:             v.Receiver,
				Content:              v.Content,
				IsRead:               v.IsRead,
				CreatedAt:            v.CreatedAt.Unix(),
			})
		}
	}
	return &message.LastMessagesResponse{Messages:messages}, nil
}