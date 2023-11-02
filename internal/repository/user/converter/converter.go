package converter

import (
	"github.com/msh2107/chat-server/internal/model"
	repoModel "github.com/msh2107/chat-server/internal/repository/user/model"
)

func ToUserFromRepo(user repoModel.User) *model.User {
	return &model.User{
		ID:     user.ID,
		ChatID: user.ChatID,
	}
}

func ToIDFromRepo(users []repoModel.User) []int64 {
	ids := make([]int64, len(users))
	for i := range users {
		ids[i] = users[i].ChatID
	}

	return ids
}
