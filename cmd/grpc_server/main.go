package main

import (
	"context"
	"fmt"
	desc "github.com/mukhinfa/chat-server/pkg/chatServer_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServerV1Server
}

func (s server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Got message with users: %v", req.GetUsername())
	return &desc.CreateResponse{}, nil
}

func (s server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Got message from: %s\nMessage: %s", req.From, req.Text)
	return &emptypb.Empty{}, nil
}

func (s server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Got request to delete chat: %d", req.Id)
	return &emptypb.Empty{}, nil
}

func main() {
	log.Println("start server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
