package models

import (
	"time"
)

type FAQ struct {
	ID        uint      `gorm:"primaryKey"`
	Question  string    `gorm:"not null"`
	Answer    string    `gorm:"not null"`
	Active    bool      `gorm:"default:true"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}