package entities

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Name    string  `gorm:"not null"`
	Address string  `gorm:"not null;unique"`
	Staff   []Staff `gorm:"many2many:business_staff"`
}

func (e *Business) Create(db *gorm.DB) error {
	err := db.Create(e).Error

	// error handling
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062: // Mysql code for DuplicatedEntry
			msg := "already exists a business with the same address"
			return errors.New(msg)
		default:
			return errors.New("couldn't create new business entity")
		}
	}
	return nil
}

func (e *Business) BeforeDelete(db *gorm.DB) (err error) {
	var found = &Business{}
	db.Find(found, *e)
	if found.CreatedAt.IsZero() {
		return errors.New("business not found")
	}
	return nil
}

func (e *Business) Delete(db *gorm.DB, id uint64) error {
	return db.Delete(e, id).Error
}
