package chat_server

import (
	"github.com/msh2107/chat-server/internal/client/db"
	"github.com/msh2107/chat-server/internal/repository"
	service "github.com/msh2107/chat-server/internal/service"
)

type serv struct {
	chatRepo    repository.ChatRepository
	messageRepo repository.MessageRepository
	userRepo    repository.UserRepository
	txManager   db.TxManager
}

func NewService(chatRepo repository.ChatRepository, messageRepo repository.MessageRepository, userRepo repository.UserRepository, txManager db.TxManager) service.ChatService {
	return &serv{chatRepo: chatRepo, messageRepo: messageRepo, userRepo: userRepo, txManager: txManager}
}
