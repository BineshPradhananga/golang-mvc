package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string
	Body        string
	Description string
	Slug        string
	UserID      uint // Foreign key referencing the User table
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
