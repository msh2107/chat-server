package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/msh2107/chat-server/internal/client/db"
	"github.com/msh2107/chat-server/internal/model"
	"github.com/msh2107/chat-server/internal/repository"
	"github.com/msh2107/chat-server/internal/repository/user/converter"
	repoModel "github.com/msh2107/chat-server/internal/repository/user/model"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user model.User) error {
	builderInsert := sq.Insert("chat_user").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id").
		Values(user.ChatID, user.ID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetChatsByUser(ctx context.Context, id int64) ([]int64, error) {
	builderSelect := sq.Select("chat_id", "user_id").From("chat_user").Where(sq.Eq{"user_id": id}).PlaceholderFormat(sq.Dollar)
	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetChatsByUser",
		QueryRaw: query,
	}
	var users []repoModel.User
	err = r.db.DB().ScanAllContext(ctx, &users, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToIDFromRepo(users), nil
}
