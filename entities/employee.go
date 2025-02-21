package entities

type Employee struct {
	User
	Businesses []Business `gorm:"many2many:business_employees"`
}
