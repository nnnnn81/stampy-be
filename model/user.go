package model

import "time"

type User struct {
	Id             uint      `gorm:"primary_key" json:"id"`
	Username       string    `gorm:"not null" json:"username"`
	Email          string    `gorm:"not null;unique" json:"email"`
	HashedPassword string    `gorm:"not null"`
	AvatarUrl      string    `gorm:"not null" json:"avatarUrl"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type OmitUser struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
}
