package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	chatImpl "github.com/msh2107/chat-server/internal/api/chat-server"
	"github.com/msh2107/chat-server/internal/model"
	"github.com/msh2107/chat-server/internal/service"
	"github.com/msh2107/chat-server/internal/service/mocks"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestGetMessages(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	type args struct {
		ctx context.Context
		req *desc.GetMessagesRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		limit = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetMessagesRequest{
			ChatId: id,
			Limit:  limit,
		}

		messages = make([]model.Message, 5)

		res = &desc.GetMessagesResponse{
			Messages: make([]*desc.Message, 0, 5),
		}
	)

	gofakeit.Slice(messages)

	for i := 0; i < len(messages); i++ {
		res.Messages = append(res.Messages, &desc.Message{
			Id: messages[i].ID,
			Info: &desc.MessageInfo{
				ChatId: messages[i].Info.ChatId,
				From:   messages[i].Info.From,
				Text:   messages[i].Info.Text,
			},
			SentAt: timestamppb.New(messages[i].SentAt),
		})
	}

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetMessagesResponse
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
				mock.GetMessagesMock.Expect(ctx, id, limit).Return(messages, nil)
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
				mock.GetMessagesMock.Expect(ctx, id, limit).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charServiceMock := tt.chatServiceMock(mc)
			api := chatImpl.NewImplementation(charServiceMock)

			empty, err := api.GetMessages(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, empty)
		})
	}
}
