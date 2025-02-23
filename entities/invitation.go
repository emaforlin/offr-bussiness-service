package entities

type Invitation struct {
	Token          string `gorm:"primaryKey"`
	InviterID      string `gorm:"not null"`
	RecipientEmail string `gorm:"not null"`
	AssignedRole   string `gorm:"not null"`
	Expiration     int64  `gorm:"not null"`
}

type InvitationDto struct {
	InviterID      string
	RecipientEmail string
	AssignedRole   string
}
