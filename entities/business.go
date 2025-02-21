package entities

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name      string     `gorm:"not null"`
	Address   string     `gorm:"not null;unique"`
	Manager   User       `gorm:"not null"`
	Employees []Employee `gorm:"many2many:business_employees"`
}
