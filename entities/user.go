package entities

type User struct {
	ID      uint   `gorm:"primaryKey"`
	Auth0ID string `gorm:"unique;not null"`
	Email   string `gorm:"unique; not null"`
}
