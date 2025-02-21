package entities

type Owner struct {
	ID         uint       `gorm:"primaryKey"`
	Auth0ID    string     `gorm:"unique;not null"`
	Email      string     `gorm:"unique;not null"`
	Businesses []Business `gorm:"many2many:business_owners"`
}
