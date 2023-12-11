package chat_server

import (
	"context"

	"github.com/msh2107/chat-server/internal/converter"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
)

func (i *Implementation) GetMessages(ctx context.Context, req *desc.GetMessagesRequest) (*desc.GetMessagesResponse, error) {
	messages, err := i.serv.GetMessages(ctx, req.GetChatId(), req.GetLimit())
	if err != nil {
		return nil, err
	}

	return &desc.GetMessagesResponse{Messages: converter.ToMessageFromService(messages)}, nil
}
