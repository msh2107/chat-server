package model

type Chat struct {
	ID   int64
	Info ChatInfo
}

type ChatInfo struct {
	OwnerID int64
	Users   []int64
}
