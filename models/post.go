package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID     string  `gorm:"primaryKey"`
	Title  string  `jsom:"title"`
	Author string  `json:"author"`
	Body   string  `json:"body"`
	Reply  []Reply `json:"replies" gorm:"constraint:onUpdate:CASCADE,OnDelete:SET NULL;"`
}
