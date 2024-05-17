package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	//lint:ignore U1000 reason
	items []Item `gorm:"constraint:OnDelete:CASCADE"` // This is a one-to-many relationship.
}
