package model

import "time"

type Notice struct {
	Id         uint      `gorm:"primary_key" json:"id"`
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Stamp      string    `json:"stamp"`
	Message    string    `json:"message"`
	CurrentDay int       `json:"currentDay"`
	IsLastDay  bool      `json:"isLastDay"`
	HrefPrefix string    `json:"hrefPrefix"`
	Sender     uint      `json:"sender"`
	Receiver   uint      `json:"receiver"`
	Read       bool      `json:"read"`
	CreatedAt  time.Time `json:"createdAt"`
	ListType   string    `json:"listType"`
	CardId     uint      `json:"cardId"`
	LetterId   uint      `json:"letterId"`
}
