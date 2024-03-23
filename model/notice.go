package model

import "time"

type Notice struct {
	Id         uint `gorm:"primary_key" json:"id"`
	Type       string
	Title      string
	Stamp      string
	Content    string
	CurrentDay int
	IsLastDay  bool
	HrefPrefix string
	Sender     uint
	Receiver   uint
	Read       bool
	CreatedAt  time.Time
	ListType   string
	CardId     uint
}
