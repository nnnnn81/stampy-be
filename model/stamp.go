package model

import (
	"time"
)

type Stamp struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	StampImg  string    `json:"stamp"`
	Message   string    `json:"message"`
	Nthday    int       `json:"nthday"`
	StampedBy uint      `json:"stampedBy"`
	CardId    uint      `json:"cardId"`
	CreatedAt time.Time `json:"createdAt"`
}
