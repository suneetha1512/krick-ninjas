package models

import "github.com/lib/pq"

type Person struct {
	Id        string         `json:"id"  gorm:"type:varchar(100);unique_index;primary_key"`
	FirstName string         `json:"firstName" gorm:"type:varchar(100)"`
	LastName  string         `json:"lastName" gorm:"type:varchar(50)"`
	Phone     string         `json:"phone" gorm:"type:varchar(15)"`
	Email     string         `json:"email" gorm:"type:varchar(30)"`
	City      string         `json:"city" gorm:"type:varchar(15)"`
	Country   string         `json:"country" gorm:"type:varchar(15)"`
	Choices   pq.StringArray `json:"choices,omitempty" gorm:"type:varchar(64)[]"`
}

type Choice struct {
	Id          string         `json:"id"  gorm:"type:varchar(100);primary_key"`
	Name        string         `json:"name" gorm:"type:varchar(50)"`
	Description string         `json:"description" gorm:"type:varchar(500)"`
	Options     pq.StringArray `json:"options,omitempty" gorm:"type:varchar(64)[]"`
}
