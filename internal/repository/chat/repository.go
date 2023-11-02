package chat

import (
	"context"
	"github.com/msh2107/chat-server/internal/client/db"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/msh2107/chat-server/internal/model"
	"github.com/msh2107/chat-server/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, chat *model.ChatInfo) (int64, error) {

	builderInsert := sq.Insert("chat").
		PlaceholderFormat(sq.Dollar).
		Columns("owner_id").
		Values(chat.OwnerID).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("builder error: %v", err.Error())
		return 0, err
	}

	var chatID int64
	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete("chat").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
