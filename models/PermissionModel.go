package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Routes       	string
	Method       	string
	UserID      	uint // Foreign key referencing the User table
	User        	User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
