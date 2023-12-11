package chat_server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/msh2107/chat-server/internal/converter"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.serv.SendMessage(ctx, converter.ToServiceFromReq(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
