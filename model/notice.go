package model

import "time"

type Notice struct {
	Id         int `gorm:"primary_key" json:"id"`
	Type       string
	Title      string
	Stamp      string
	Content    string
	HrefPrefix string
	Sender     User
	Receiver   User
	Read       bool
	CreatedAt  time.Time
	SendAt     string
	ListType   string
}
