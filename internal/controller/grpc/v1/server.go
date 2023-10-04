package v1

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/msh2107/chat-server/pkg/chat_v1"
)

// ChatServer - .
type ChatServer struct {
	desc.UnimplementedChatV1Server
}

// NewChatServer - .
func NewChatServer() *ChatServer {
	return &ChatServer{}
}

// Create - .
func (s *ChatServer) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Recieved Create:\n\t Usernames: %v\n", req.GetUsernames())
	randInt64, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<62))
	if err != nil {
		return nil, err
	}

	id := randInt64.Int64()

	return &desc.CreateResponse{Id: id}, nil
}

// Delete - .
func (s *ChatServer) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Recieved Delete:\n\t ID: %d \n", req.GetId())
	return &emptypb.Empty{}, nil
}

// SendMessage - .
func (s *ChatServer) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Reciev SendMessage:\n\t From: %s\n\t Text: %s\n\t Timestamp: %v\n", req.GetFrom(), req.GetText(), req.GetTimestamp().AsTime())

	return &emptypb.Empty{}, nil
}
