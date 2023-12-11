package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	chatImpl "github.com/msh2107/chat-server/internal/api/chat-server"
	"github.com/msh2107/chat-server/internal/service"
	"github.com/msh2107/chat-server/internal/service/mocks"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestDelete(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)
		id  = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &desc.DeleteRequest{
			ChatId: id,
		}

		res = &emptypb.Empty{}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
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
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
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
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charServiceMock := tt.chatServiceMock(mc)
			api := chatImpl.NewImplementation(charServiceMock)

			empty, err := api.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, empty)
		})
	}
}
