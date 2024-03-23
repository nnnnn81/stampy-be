package model

import "time"

type Stampcard struct {
	Id            uint      `gorm:"primary_key" json:"id"`
	Title         string    `gorm:"not null" json:"title"`
	CreatedBy     uint      `json:"CreatedBy"`
	JoinedUser    uint      `json:"JoinedUser"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	StartDate     string    `json:"startDate"`
	EndDate       string    `json:"endDate"`
	Days          int
	CurrentDay    int     `json:"CurrentDay"`
	IsStampy      bool    `json:"isStampy"`
	IsCompleted   bool    `json:"IsCompleted"`
	IsDeleted     bool    `json:"IsDeleted"`
	StampNodes    []Stamp `gorm:"foreignKey:CardId" json:"stampNodes"`
	BackgroundUrl string  `json:"BackgroundUrl"`
}
