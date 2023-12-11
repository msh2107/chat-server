package chat_server

import (
	"context"
	"errors"
	"fmt"

	"github.com/msh2107/chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, message *model.MessageInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		chats, errTx := s.userRepo.GetChatsByUser(ctx, message.From)
		if errTx != nil {
			return errTx
		}
		fmt.Println(chats)
		if !in(chats, message.ChatId) {
			return errors.New("this user not in the chat")
		}
		errTx = s.messageRepo.SendMessage(ctx, message)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func in(arr []int64, t int64) bool {
	for _, i := range arr {
		if i == t {
			return true
		}
	}
	return false
}
