package models

type TeamMember struct {
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"not null"`
	Designation       string `gorm:"not null"`
	Email             string `gorm:"size:150;unique"`
	ProfilePictureURL string `gorm:"size:255"`
}
