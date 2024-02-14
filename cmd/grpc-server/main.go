package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/sarastee/chat-server/pkg/chat_api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	grpcPort = 50052
)

type server struct {
	desc.UnimplementedChatAPIV1Server
}

type chat struct {
	ID        int64
	Usernames []string
}

var testChat chat

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	testChat.ID = 0
	testChat.Usernames = req.Usernames

	log.Printf("Chat has been created: %+v", testChat)

	return &desc.CreateResponse{
		Id: testChat.ID,
	}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if req.Id != testChat.ID {
		return nil, fmt.Errorf("chat %d not found", req.Id)
	}

	testChat = chat{}

	log.Printf("Chat %d was deleted", req.Id)

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Message: \"%v\" from: %v at %v", req.Message.Text, req.Message.From, req.Message.Timestamp.AsTime())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatAPIV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
