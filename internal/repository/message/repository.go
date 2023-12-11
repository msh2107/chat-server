package message

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/msh2107/chat-server/internal/client/db"
	"github.com/msh2107/chat-server/internal/model"
	"github.com/msh2107/chat-server/internal/repository"
	"github.com/msh2107/chat-server/internal/repository/message/converter"
	repoModel "github.com/msh2107/chat-server/internal/repository/message/model"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.MessageRepository {
	return &repo{db: db}
}

func (r *repo) SendMessage(ctx context.Context, message *model.MessageInfo) error {
	builderInsert := sq.Insert("message").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text").
		Values(message.ChatId, message.From, message.Text)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetMessages(ctx context.Context, chatID int64, limit int64) ([]model.Message, error) {
	builderSelect := sq.Select("id", "sent_at", "chat_id", "user_id", "text").
		From("message").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": chatID}).
		Limit(uint64(limit))

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "chat_repository.GetMessages",
		QueryRaw: query,
	}

	var messages []repoModel.Message
	err = r.db.DB().ScanAllContext(ctx, &messages, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToMessageFromRepo(messages), nil
}
