package chat_server

import (
	"context"
	"github.com/msh2107/chat-server/internal/model"
)

func (s *serv) GetMessages(ctx context.Context, chatID int64, limit int64) ([]model.Message, error) {
	messages, err := s.messageRepo.GetMessages(ctx, chatID, limit)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
