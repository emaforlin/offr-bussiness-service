package repository

import "gorm.io/gorm"

// Gorm Managed Database
type GMDatabase interface {
	Connect() error
	Cursor() *gorm.DB
}
