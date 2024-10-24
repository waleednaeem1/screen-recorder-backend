package models

import (
	"time"
)

type ProfileType struct {
	ID                   uint      `gorm:"primaryKey"`
	RecordingsLeft       int       `gorm:"default:10"`
	VideoLengthAllocated int       `gorm:"default:5"`
	Type                 string    `json:"type" gorm:"not null"`
	CreatedAt            time.Time `gorm:"autoCreateTime"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime"`
}
