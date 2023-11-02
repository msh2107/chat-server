package chat_server

import (
	"github.com/msh2107/chat-server/internal/service"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	serv service.ChatService
}

func NewImplementation(serv service.ChatService) *Implementation {
	return &Implementation{serv: serv}
}
