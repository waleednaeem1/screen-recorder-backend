package models

import (
	"time"
)

type PrivacyPolicy struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
