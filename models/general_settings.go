package models

import (
	"time"
)

type GeneralSettings struct {
	ID            uint      `gorm:"primaryKey"`
	SupportEmail  string    `gorm:"size:150"`
	SupportNumber string    `gorm:"size:150"`
	SocialLinks   string    `gorm:"type:text"`
	Logo          string    `gorm:"type:text"`
	Favicon       string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

