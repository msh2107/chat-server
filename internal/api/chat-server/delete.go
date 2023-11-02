package chat_server

import (
	"context"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.serv.Delete(ctx, req.GetChatId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
