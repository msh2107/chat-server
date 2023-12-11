package service

import (
	"context"

	"github.com/msh2107/chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message *model.MessageInfo) error
	GetMessages(ctx context.Context, chatID int64, limit int64) ([]model.Message, error)
}
