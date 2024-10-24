package models

import (
	"time"
)

type Pricing struct {
	ID        uint   `gorm:"primaryKey"`
	PlanName  string `gorm:"not null"`
	Price     float64
	Features  string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
