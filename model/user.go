package model

import "time"

type User struct {
	Id             int       `gorm:"primary_key" json:"id"`
	Username       string    `gorm:"not null" json:"username"`
	Email          string    `gorm:"not null;unique" json:"email"`
	HashedPassword string    `gorm:"not null"`
	AvaterUrl      string    `gorm:"not null" json:"avaterUrl"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
