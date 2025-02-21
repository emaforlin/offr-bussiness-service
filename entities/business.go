package entities

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name    string  `gorm:"not null"`
	Address string  `gorm:"not null;unique"`
	Staff   []Staff `gorm:"many2many:business_staff"`
}
