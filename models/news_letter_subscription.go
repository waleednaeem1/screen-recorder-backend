package models

import (
	"time"
)

type NewsletterSubscription struct {
	ID           uint      `gorm:"primaryKey"`
	Email        string    ` json:"email"`
	SubscribedAt time.Time `json:"subscribed_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
