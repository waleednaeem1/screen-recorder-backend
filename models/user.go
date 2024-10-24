package models

import (
	"time"
)

type User struct {
	ID               uint              `gorm:"primaryKey"`
	Name             string            `gorm:"size:150"`
	Email            string            `gorm:"size:150;unique;not null"`
	CompanyName      string            `gorm:"size:100"`
	Password         string            `gorm:"not null"`
	Token            string            `gorm:"size:255"`
	UserType         string            `json:"user_type" gorm:"not null"`
	ProfileTypeID    *uint             `json:"profile_type_id"`
	ProfileType      ProfileType       `gorm:"foreignKey:ProfileTypeID;references:ID"`
	ScreenRecordings []ScreenRecording `gorm:"foreignKey:UserID"`
	Status           *UserStatus       `json:"status"`
	CreatedAt        time.Time         `gorm:"autoCreateTime"`
	UpdatedAt        time.Time         `gorm:"autoUpdateTime"`
}

type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
)
