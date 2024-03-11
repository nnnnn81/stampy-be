package model

import "time"

type Stamp struct {
	Id        int       `gorm:"primary_key" json:"id"`
	StampImg  string    `json:"stamp"`
	Message   string    `json:"message"`
	Stamped   bool      `json:"stamped"`
	Nthday    int       `json:"nthday"`
	StampedBy User      `json:"stampedBy"`
	Card      int       `json:"cardId"`
	CreatedAt time.Time `json:"createdAt"`
}
