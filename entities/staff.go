package entities

type Staff struct {
	ID         uint       `gorm:"primaryKey"`
	Auth0ID    string     `gorm:"not null;unique"`
	Businesses []Business `gorm:"many2many:business_staff"`
}
