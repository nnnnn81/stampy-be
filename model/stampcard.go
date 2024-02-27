package model

import "time"

type Stampcard struct {
	Id int `gorm: "primary_key" json:"id"`
	Title string `gorm: "not null" json:"title"`
	CreatedBy User `json:"CreatedBy"`
	JoinedUser User `json:"JoinedUser"`
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
	CurrentDay int `json:"CurrentDay"`
	IsCompleted bool `json:"IsCompleted"`
	IsDeleted bool `json:"IsDeleted"`
	StampNodes []Stamp `json:"stampNodes"`
	BackgroundUrl string `json:"BackgroundUrl"`

}