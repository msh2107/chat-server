package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/msh2107/chat-server/config"
	v1 "github.com/msh2107/chat-server/internal/controller/grpc/v1"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	address := fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.GRPCPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chatServer := v1.NewChatServer()
	desc.RegisterChatV1Server(s, chatServer)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("chat server listening on %s", address)

	<-interrupt

	log.Printf("Shutting down...")
	s.GracefulStop()

}
