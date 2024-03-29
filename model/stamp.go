package model

import (
	"time"
)

type Stamp struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	StampImg  string    `json:"stamp"`
	Message   string    `json:"message"`
	NthDay    int       `json:"nthday"`
	StampedBy uint      `json:"stampedBy"`
	StampedTo uint      `json:"stampedTo"`
	Stamped   bool      `json:"stamped"`
	CardId    uint      `json:"cardId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
