package model

import "time"

type Stampcard struct {
	Id            uint      `gorm:"primary_key" json:"id"`
	Title         string    `gorm:"not null" json:"title"`
	CreatedBy     uint      `json:"createdBy"`
	JoinedUser    uint      `json:"joinedUser"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	StartDate     string    `json:"startDate"`
	EndDate       string    `json:"endDate"`
	Days          int       `json:"days"`
	CurrentDay    int       `json:"currentDay"`
	IsStampy      bool      `json:"isStampy"`
	IsCompleted   bool      `json:"isCompleted"`
	IsDeleted     bool      `json:"isDeleted"`
	StampNodes    []Stamp   `gorm:"foreignKey:CardId" json:"stampNodes"`
	LetterId      uint      `json:"letterId"`
	BackgroundUrl string    `json:"backgroundUrl"`
}
