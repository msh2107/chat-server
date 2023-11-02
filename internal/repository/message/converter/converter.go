package converter

import (
	"github.com/msh2107/chat-server/internal/model"
	repoModel "github.com/msh2107/chat-server/internal/repository/message/model"
)

func ToMessageFromRepo(messages []repoModel.Message) []model.Message {
	convertedMessages := make([]model.Message, 0, len(messages))
	for _, message := range messages {
		convertedMessages = append(convertedMessages, model.Message{
			ID:     message.ID,
			Info:   ToMessageInfoFromRepo(message.Info),
			SentAt: message.SentAt,
		})
	}

	return convertedMessages
}

func ToMessageInfoFromRepo(info repoModel.MessageInfo) model.MessageInfo {
	return model.MessageInfo{
		ChatId: info.ChatId,
		From:   info.From,
		Text:   info.Text,
	}
}
