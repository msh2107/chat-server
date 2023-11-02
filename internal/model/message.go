package model

import "time"

type Message struct {
	ID     int64
	Info   MessageInfo
	SentAt time.Time
}

type MessageInfo struct {
	ChatId int64
	From   int64
	Text   string
}
