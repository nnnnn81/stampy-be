package model

import "time"

type Stampcard struct {
	Id            uint      `gorm:"primary_key" json:"id"`
	Title         string    `gorm:"not null" json:"title"`
	CreatedBy     uint      `json:"CreatedBy"`
	JoinedUser    string    `json:"JoinedUser"`
	Created_at    time.Time `json:"createdAt"`
	Updated_at    time.Time `json:"updatedAt"`
	CurrentDay    int       `json:"CurrentDay"`
	IsStampy      bool      `json:"isStampy"`
	IsCompleted   bool      `json:"IsCompleted"`
	IsDeleted     bool      `json:"IsDeleted"`
	BackgroundUrl string    `json:"BackgroundUrl"`
}
