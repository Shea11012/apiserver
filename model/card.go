package model

import (
	"time"
)

type Card struct {
	Id 	uint64	`gorm:"primary_key;auto_increment;column:id" json:"-"`
	UserId uint64 `gorm:"column:userid" json:"userId"`
	Name string `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	DeletedAt *time.Time `gorm:"column:deletedAt" json:"deletedAt"`
}


func (c *Card) Create() error {
	return DB.mysql.Create(&c).Error
}

func (c *Card) Get() ([]*Card,error) {
	cards := make([]*Card,0)
	if err := DB.mysql.Where("userid = ?",c.UserId).Find(&cards).Error; err != nil {
		return cards,err
	}
	return cards,nil
}