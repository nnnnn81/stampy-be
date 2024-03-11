package model

import "time"

type Stamp struct {
	Id        int       `gorm:"primary_key" json:"id"`
	StampImg  string    `json:"stamp"`
	Message   string    `json:"message"`
	Stamped   bool      `json:"stamped"`
	stampedAt time.Time `json:"stampedAt"`
	Nthday    int       `json:"nthday"`
	StampedBy User      `json:"stampedBy"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
	CardId    int       `gorm: "not null" json:"cardId"`
	CreatedAt time.Time `json:"createdAt"`
}
