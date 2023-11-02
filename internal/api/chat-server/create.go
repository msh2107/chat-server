package chat_server

import (
	"context"
	"github.com/msh2107/chat-server/internal/converter"
	"github.com/msh2107/chat-server/internal/model"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chat := model.ChatInfo{OwnerID: req.GetOwnerId(), Users: converter.ToUsersFromReq(req.GetUsers())}
	id, err := i.serv.Create(ctx, &chat)
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{ChatId: id}, nil
}
