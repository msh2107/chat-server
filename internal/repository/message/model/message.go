package model

import "time"

type Message struct {
	ID     int64       `db:"id"`
	Info   MessageInfo `db:""`
	SentAt time.Time   `db:"sent_at"`
}

type MessageInfo struct {
	ChatId int64  `db:"chat_id"`
	From   int64  `db:"user_id"`
	Text   string `db:"text"`
}
