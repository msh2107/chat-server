package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/msh2107/chat-server/internal/model"
	desc "github.com/msh2107/chat-server/pkg/chat_v1"
)

func ToServiceFromCreateReq(req *desc.CreateRequest) *model.ChatInfo {
	chat := model.ChatInfo{
		OwnerID: req.GetOwnerId(),
		Users:   ToUsersFromReq(req.GetUsers()),
	}
	return &chat
}

func ToMessageFromService(messages []model.Message) []*desc.Message {
	convertedMessages := make([]*desc.Message, 0, len(messages))
	for _, message := range messages {
		convertedMessages = append(convertedMessages, &desc.Message{
			Id:     message.ID,
			Info:   ToMessageInfoFromService(message.Info),
			SentAt: timestamppb.New(message.SentAt),
		})
	}

	return convertedMessages
}

func ToMessageInfoFromService(info model.MessageInfo) *desc.MessageInfo {
	return &desc.MessageInfo{
		ChatId: info.ChatId,
		From:   info.From,
		Text:   info.Text,
	}
}

func ToServiceFromReq(info *desc.MessageInfo) *model.MessageInfo {
	return &model.MessageInfo{
		ChatId: info.GetChatId(),
		From:   info.GetFrom(),
		Text:   info.GetText(),
	}
}

func ToUsersFromReq(ids []int64) []model.User {
	users := make([]model.User, len(ids))
	for i := range ids {
		users[i].ID = ids[i]
	}
	return users
}
