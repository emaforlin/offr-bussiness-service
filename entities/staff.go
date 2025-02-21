package entities

type Staff struct {
	ID         uint       `gorm:"primaryKey"`
	Businesses []Business `gorm:"many2many:business_staff"`
}
