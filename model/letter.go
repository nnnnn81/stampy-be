package model

import "time"

type Letter struct {
	Id         uint      `gorm:"primary_key" json:"id"`
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Stamp      string    `json:"stamp"`
	Content    string    `json:"content"`
	HrefPrefix string    `json:"hrefPrefix"`
	Sender     uint      `json:"sender"`
	Receiver   uint      `json:"receiver"`
	Read       bool      `json:"read"`
	IsVisible  bool      `json:"isVisible"`
	CreatedAt  time.Time `json:"createdAt"`
	ListType   string    `json:"listType"`
	CardId     uint      `json:"cardId"`
}
