package chat_server

import (
	"context"

	"github.com/msh2107/chat-server/internal/converter"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.serv.Create(ctx, converter.ToServiceFromCreateReq(req))
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{ChatId: id}, nil
}
