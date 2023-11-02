package model

type User struct {
	ID     int64 `db:"user_id"`
	ChatID int64 `db:"chat_id"`
}
