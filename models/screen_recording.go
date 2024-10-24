package models

import "time"

type ScreenRecording struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	VideoURL    string    `gorm:"not null"`
	VideoTitle  string    `gorm:"size:100"`
	VideoLength int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
