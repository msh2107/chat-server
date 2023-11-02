package repository

import (
	"context"

	"github.com/msh2107/chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	GetChatsByUser(ctx context.Context, id int64) ([]int64, error)
}

type MessageRepository interface {
	SendMessage(ctx context.Context, message *model.MessageInfo) error
	GetMessages(ctx context.Context, chatID int64, limit int64) ([]model.Message, error)
}
