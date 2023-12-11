package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	chatImpl "github.com/msh2107/chat-server/internal/api/chat-server"
	"github.com/msh2107/chat-server/internal/model"
	"github.com/msh2107/chat-server/internal/service"
	serviceMocks "github.com/msh2107/chat-server/internal/service/mocks"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)
		ids = make([]int64, 5)
		id  = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Users:   ids,
			OwnerId: id,
		}

		users = make([]model.User, len(ids))

		info = &model.ChatInfo{
			OwnerID: id,
			Users:   users,
		}

		res = &desc.CreateResponse{ChatId: id}
	)

	gofakeit.Slice(&ids)
	for i := 0; i < len(users); i++ {
		users[i].ID = ids[i]
	}

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charServiceMock := tt.chatServiceMock(mc)
			api := chatImpl.NewImplementation(charServiceMock)

			id, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, id)
		})
	}
}
