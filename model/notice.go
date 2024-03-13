package model

import "time"

type Notice struct {
	Id         uint `gorm:"primary_key" json:"id"`
	Type       string
	Title      string
	Stamp      string
	Content    string
	HrefPrefix string
	Sender     string
	Receiver   string
	Read       bool
	CreatedAt  time.Time
	ListType   string
}
