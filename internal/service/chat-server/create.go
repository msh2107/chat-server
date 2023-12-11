package chat_server

import (
	"context"

	"github.com/msh2107/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, chat *model.ChatInfo) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepo.Create(ctx, chat)
		if errTx != nil {
			return errTx
		}
		for _, user := range chat.Users {
			errTx = s.userRepo.Create(ctx, model.User{
				ID:     user.ID,
				ChatID: id,
			})
			if errTx != nil {
				return errTx
			}
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
